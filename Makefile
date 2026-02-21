RUNTIME    := $(shell which podman 2>/dev/null || which docker)
BINARY     := pgxray
CMD        := ./src/cmd/pgxray
BUILD_DIR  := ./bin
IMAGE_DEV  := pgxray-dev
IMAGE_LINT := pgxray-lint

.PHONY: all image-dev image-lint build run test lint tidy clean

all: build

image-dev:
	$(RUNTIME) build -f container/dev/Containerfile -t $(IMAGE_DEV) .

image-lint: image-dev
	$(RUNTIME) build -f container/lint/Containerfile -t $(IMAGE_LINT) .

tidy:
	go mod tidy

build: image-dev
	$(RUNTIME) run --rm -v $(PWD)/$(BUILD_DIR):/out $(IMAGE_DEV) \
		go build -o /out/$(BINARY) $(CMD)

test: image-dev
	$(RUNTIME) run --rm $(IMAGE_DEV) go test ./src/...

lint: image-lint
	$(RUNTIME) run --rm $(IMAGE_LINT) golangci-lint run ./src/...

run:
	./$(BUILD_DIR)/$(BINARY)

start: build run

clean:
	rm -rf $(BUILD_DIR)
