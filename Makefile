# Go parameters
go := go
go_flags :=

go_test := $(go) test
binary_name := $(shell basename $(CURDIR))
binary_name_lower :=  $(shell echo $(binary_name) | tr '[:upper:]' '[:lower:]')
build_dir := ./build
src_files := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Default target
.PHONY: all
all: build clean docker

# Build target
.PHONY: build
build: $(build_dir)/$(binary_name)

$(build_dir)/$(binary_name): $(src_files)
	@mkdir -p $(build_dir)
	$(go) build $(go_flags) -o $(build_dir)/$(binary_name)

# Clean target
.PHONY: clean
clean:
	@rm -rf $(build_dir)

# Test target
.PHONY: test
test:
	$(go_test) -v ./...

# Create docker image
.PHONY: docker
docker:
	@RELEASE_TAG=$$(git describe --tags 2>/dev/null); \
	if [ -n "$$RELEASE_TAG" ]; then \
		docker build -t $(binary_name_lower):$$RELEASE_TAG . ; \
	else \
		docker build -t $(binary_name_lower):dev . ; \
	fi

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build    - Build the binary"
	@echo "  clean    - Clean build artifacts"
	@echo "  docker   - Create Docker image"
	@echo "  help     - Show this help message"
