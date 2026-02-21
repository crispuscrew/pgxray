RUNTIME   := $(shell which podman 2>/dev/null || which docker)
BINARY    := pgxray
CMD       := ./cmd/pgxray
BUILD_DIR := ./bin
SRC       := src
IMAGE_DEV  := pgxray-dev
IMAGE_LINT := pgxray-lint

.PHONY: all image-dev image-lint build run test lint tidy clean

all: build

image-dev:
	$(RUNTIME) build -f container/dev/Containerfile -t $(IMAGE_DEV) .

image-lint: image-dev
	$(RUNTIME) build -f container/lint/Containerfile -t $(IMAGE_LINT) .

tidy: image-dev
	$(RUNTIME) run --rm $(IMAGE_DEV) go mod tidy

build: image-dev
	$(RUNTIME) run --rm -v $(PWD)/$(BUILD_DIR):/out $(IMAGE_DEV) \
		go build -o /out/$(BINARY) $(CMD)

test: image-dev
	$(RUNTIME) run --rm $(IMAGE_DEV) go test ./...

lint: image-lint
	$(RUNTIME) run --rm $(IMAGE_LINT) golangci-lint run ./...

run:
	./$(BUILD_DIR)/$(BINARY)

start: build run

clean:
	rm -rf $(BUILD_DIR)
