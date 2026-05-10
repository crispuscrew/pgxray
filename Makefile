BOLD  := \033[1m
CYAN  := \033[36m
RESET := \033[0m

RUNTIME    := $(shell which podman 2>/dev/null || which docker)
BINARY     := pgxray
CMD        := ./cmd/pgxray
BUILD_DIR  := ./bin
IMAGE_DEV  := pgxray-dev
IMAGE_LINT := pgxray-lint
IMAGE_TEST := pgxray-test
RUN_DEV    := $(RUNTIME) run --rm
TEST_NET   := pgxray-test-net
TEST_PG    := pgxray-test-pg
TEST_PKG   ?= ./...
TEST_PG_ARGS := \
	-e POSTGRES_PASSWORD=test \
	-e POSTGRES_DB=testdb \
	-v $(PWD)/src/internal/db/testdata/seed.sql:/docker-entrypoint-initdb.d/seed.sql:ro,z \
	docker.io/library/postgres:16-alpine

.PHONY: build run lint tidy clean test db help rebuild

_image-dev:
	@$(RUNTIME) image exists $(IMAGE_DEV) || \
		$(RUNTIME) build -f container/dev/Containerfile -t $(IMAGE_DEV) .

_image-lint: _image-dev
	@$(RUNTIME) image exists $(IMAGE_LINT) || \
		$(RUNTIME) build -f container/lint/Containerfile -t $(IMAGE_LINT) .

_image-test: _image-dev
	@$(RUNTIME) image exists $(IMAGE_TEST) || \
		$(RUNTIME) build -f container/test/Containerfile -t $(IMAGE_TEST) .

build: _image-dev ## Build binary into ./bin/
	@mkdir -p $(BUILD_DIR)
	$(RUN_DEV) -v $(PWD)/src:/app -v $(PWD)/$(BUILD_DIR):/out $(IMAGE_DEV) \
		go build -o /out/$(BINARY) $(CMD)

run: ## Run compiled binary (needs live terminal)
	./$(BUILD_DIR)/$(BINARY)

tidy: _image-dev ## Run go mod tidy
	$(RUN_DEV) -v $(PWD)/src:/app $(IMAGE_DEV) go mod tidy

test: _image-test ## Run tests against containerized postgres
	$(RUNTIME) network create $(TEST_NET) 2>/dev/null ; \
	$(RUNTIME) run -d --name $(TEST_PG) --network $(TEST_NET) $(TEST_PG_ARGS) ; \
	until $(RUNTIME) exec $(TEST_PG) pg_isready -U postgres; do sleep 1; done ; \
	$(RUNTIME) run --rm --network $(TEST_NET) \
		-v $(PWD)/src:/app \
		-e PGPASSWORD=test \
		$(IMAGE_TEST) go test -v $(TEST_PKG) ; \
	$(RUNTIME) rm -f $(TEST_PG) 2>/dev/null ; \
	$(RUNTIME) network rm $(TEST_NET) 2>/dev/null

db: ## Start postgres on :5432 for manual testing
	$(RUNTIME) run --rm --name $(TEST_PG) -p 5432:5432 $(TEST_PG_ARGS)

lint: _image-lint ## Run golangci-lint
	$(RUN_DEV) -v $(PWD)/src:/app $(IMAGE_LINT) golangci-lint run ./...

clean: ## Remove build artifacts
	rm -rf $(BUILD_DIR)

rebuild: ## Force-rebuild all container images
	$(RUNTIME) rmi -f $(IMAGE_DEV) $(IMAGE_LINT) $(IMAGE_TEST) 2>/dev/null ; \
	$(RUNTIME) build -f container/dev/Containerfile -t $(IMAGE_DEV) . && \
	$(RUNTIME) build -f container/lint/Containerfile -t $(IMAGE_LINT) . && \
	$(RUNTIME) build -f container/test/Containerfile -t $(IMAGE_TEST) .

help: ## Show this help
	@printf "\n$(BOLD)pgxray$(RESET)\n\n"
	@grep -hE '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  $(CYAN)%-20s$(RESET) %s\n", $$1, $$2}'
	@printf "\n$(BOLD)Variables:$(RESET)\n"
	@printf "  $(CYAN)%-20s$(RESET) %s\n" "TEST_PKG=<pattern>" "go test package filter (default: ./...)"
	@printf "  $(CYAN)%-20s$(RESET) %s\n" "RUNTIME=<path>"     "container runtime override (default: podman or docker)"
	@printf "\n"
