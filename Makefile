RUNTIME    := $(shell which podman 2>/dev/null || which docker)
BINARY     := pgxray
CMD        := ./src/cmd/pgxray
BUILD_DIR  := ./bin
IMAGE_DEV  := pgxray-dev
IMAGE_LINT := pgxray-lint
RUN_DEV    := $(RUNTIME) run --rm

.PHONY: all build run start test lint tidy clean image-dev image-lint

all: build

image-dev:
	$(RUNTIME) build -f container/dev/Containerfile -t $(IMAGE_DEV) .

image-lint: image-dev
	$(RUNTIME) build -f container/lint/Containerfile -t $(IMAGE_LINT) .

tidy: image-dev
	$(RUN_DEV) -v $(PWD)/go.mod:/app/go.mod -v $(PWD)/go.sum:/app/go.sum \
		$(IMAGE_DEV) go mod tidy

build: image-dev
	$(RUN_DEV) -v $(PWD)/$(BUILD_DIR):/out $(IMAGE_DEV) \
		go build -o /out/$(BINARY) $(CMD)

test: image-dev
	$(RUN_DEV) $(IMAGE_DEV) go test ./src/...

lint: image-lint
	$(RUN_DEV) $(IMAGE_LINT) golangci-lint run ./src/...

run:
	./$(BUILD_DIR)/$(BINARY)

start: tidy build run

clean:
	rm -rf $(BUILD_DIR)
