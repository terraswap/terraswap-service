#
# terraswap-service
#
# build:
#   docker build --force-rm -t REPO/terraswap-service .
# run:
#   docker run --rm -it --env-file=path/to/.env --name terraswap-service -p 80:8080 REPO/terraswap-service

### BUILD
FROM golang:1.23-alpine AS build
WORKDIR /app

# Create appuser.
RUN adduser -D -g '' appuser
# Install required binaries
RUN apk add --update --no-cache zip git make build-base

# Copy app dependencies
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile
# Download all golang package dependencies
RUN make deps

# Copy source files
COPY . .

# See https://github.com/CosmWasm/wasmvm/releases
RUN set -eux; \
    export ARCH=$(uname -m); \
    WASM_VERSION=$(go list -mod=readonly -m all | grep github.com/CosmWasm/wasmvm | awk '{print $2}'); \
    if [ ! -z "${WASM_VERSION}" ]; then \
      wget -O /lib/libwasmvm_muslc.a https://github.com/CosmWasm/wasmvm/releases/download/${WASM_VERSION}/libwasmvm_muslc.${ARCH}.a; \
    fi;

# Build executable
## force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN go build -mod=readonly -tags "netgo muslc" -ldflags '-X "github.com/cosmos/cosmos-sdk/version.BuildTags=netgo,muslc" -w -s' -trimpath -o ./main ./cmd/api

### RELEASE
FROM alpine:latest AS release
RUN apk add --update --no-cache gcc

WORKDIR /app
# Expose application port
# Import the user and group files to run the app as an unpriviledged user
COPY --from=build /etc/passwd /etc/passwd

# Use an unprivileged user
USER appuser
COPY --from=build /app/cmd /app/cmd
# Grab compiled binary from build
COPY --from=build /app/main /app/main

# Expose application port
ENV APP_PORT=8000
EXPOSE $APP_PORT
# Set entry point
CMD [ "./main" ]
