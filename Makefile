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

.PHONY: build run lint tidy clean image-dev image-lint test db help rebuild

_image-dev:
	@$(RUNTIME) image exists $(IMAGE_DEV) || \
		$(RUNTIME) build -f container/dev/Containerfile -t $(IMAGE_DEV) .

_image-lint: _image-dev
	@$(RUNTIME) image exists $(IMAGE_LINT) || \
		$(RUNTIME) build -f container/lint/Containerfile -t $(IMAGE_LINT) .

_image-test: _image-dev
	@$(RUNTIME) image exists $(IMAGE_TEST) || \
		$(RUNTIME) build -f container/test/Containerfile -t $(IMAGE_TEST) .

rebuild:
	$(RUNTIME) rmi -f $(IMAGE_DEV) $(IMAGE_LINT) $(IMAGE_TEST) 2>/dev/null ; \
	$(RUNTIME) build -f container/dev/Containerfile -t $(IMAGE_DEV) . && \
	$(RUNTIME) build -f container/lint/Containerfile -t $(IMAGE_LINT) . && \
	$(RUNTIME) build -f container/test/Containerfile -t $(IMAGE_TEST) .

tidy: _image-dev
	$(RUN_DEV) -v $(PWD)/src:/app $(IMAGE_DEV) go mod tidy

build: _image-dev
	@mkdir -p $(BUILD_DIR)
	$(RUN_DEV) -v $(PWD)/src:/app -v $(PWD)/$(BUILD_DIR):/out $(IMAGE_DEV) \
		go build -o /out/$(BINARY) $(CMD)


test: _image-test
	$(RUNTIME) network create $(TEST_NET) 2>/dev/null ; \
	$(RUNTIME) run -d --name $(TEST_PG) --network $(TEST_NET) $(TEST_PG_ARGS) ; \
	until $(RUNTIME) exec $(TEST_PG) pg_isready -U postgres; do sleep 1; done ; \
	$(RUNTIME) run --rm --network $(TEST_NET) \
		-v $(PWD)/src:/app \
		-e PGPASSWORD=test \
		$(IMAGE_TEST) go test -v $(TEST_PKG) ; \
	$(RUNTIME) rm -f $(TEST_PG) 2>/dev/null ; \
	$(RUNTIME) network rm $(TEST_NET) 2>/dev/null

db:
	$(RUNTIME) run --rm --name $(TEST_PG) -p 5432:5432 $(TEST_PG_ARGS)

lint: _image-lint
	$(RUN_DEV) -v $(PWD)/src:/app $(IMAGE_LINT) golangci-lint run ./...

run:
	./$(BUILD_DIR)/$(BINARY)

clean:
	rm -rf $(BUILD_DIR)

help:
	@cat FOR_DEVELOP.md
