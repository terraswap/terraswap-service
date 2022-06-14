# TerraSwap Service
---

## Getting Started

### Requirements

- [Go installation](https://golang.org/dl/) (preferably v1.18) and a [correctly configured path](https://golang.org/doc/install#install).
- A local [docker, docker-compose](https://docs.docker.com) 
- [Golangci-lint](https://github.com/golangci/golangci-lint) to improve code quality


### Quick start

```zsh
git clone https://github.com/terraswap/terraswap-service
cd terraswap-service
vim config.yaml # to change environment variables.
make up 
make watch # server is automatically restarted when code is edited
# ...
make down # shut down all services
```

## Commands

### Run

```zsh
make all          # Same as make test
make deps         # Download dependencies
make watch        # Run development server, recompile and run tests on file change
make clean        # Remove compiled binaries from local disk
make fmt          # Format code
make lint         # Run golangci-lint on code changed since forked from master branch
make prune-deps   # Remove unused dependencies from go.mod & go.sum
make image        # Create docker image with minimal binary
make build-docker # Build with special params to create a complete binary,
                  # see Dockerfile

```

### Test

```zsh
make test       # Run tests for all packages
make cover      # Check coverage for all packages
```

## Packages

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/terraswap/terraswap-service/issues) to report any bugs or ask feature requests.

## License
