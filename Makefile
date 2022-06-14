.PHONY: all
all: test

# Start the minimum requirements for the service, i.e. db
.PHONY: up
up:
	docker-compose up -d

# Stop all services
.PHONY: down
down:
	docker-compose down

# Explicitly install dependencies. In most cases this is not required as go will automatically download missing deps.
.PHONY: deps
deps:
	go mod download

# Build the main executable
main:
	go build -o main ./cmd/api

# This is a specialized build for running the executable inside a minimal scratch container
.PHONY: build-api
build-api:
	go build -ldflags="-w -s" -a -o ./main ./cmd/api

# Watch for source code changes to recompile + test
.PHONY: watch
watch:
	GO111MODULE=off go get github.com/cortesi/modd/cmd/modd
	modd

# Run all unit tests
.PHONY: test
test: main
	go test -short ./...

# Run all benchmarks
.PHONY: bench
bench:
	go test -short -bench=. ./...

# Same as test but with coverage turned on
.PHONY: cover
cover:
	go test -short -cover -covermode=atomic ./...


# Apply https://golang.org/cmd/gofmt/ to all packages
.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: fmt-check
fmt-check:
ifneq ($(shell gofmt -l .),)
	$(error gofmt fail in $(shell gofmt -l .))
endif

# Apply https://github.com/golangci/golangci-lint to changes since forked from master branch
.PHONY: lint
lint:
	golangci-lint run --timeout=5m --new-from-rev=$(shell git merge-base $(shell git branch | sed -n -e 's/^\* \(.*\)/\1/p') github/main) --enable=unparam --enable=misspell --enable=prealloc

# Remove all compiled binaries from the directory
.PHONY: clean
clean:
	go clean

# Analyze the code for any unused dependencies
.PHONY: prune-deps
prune-deps:
	go mod tidy

# Create the service docker image
.PHONY: image
image:
	docker build --force-rm -t terraswap/terraswap-service .
