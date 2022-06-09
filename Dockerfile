#
# terraswap-service
#
# build:
#   docker build --force-rm -t REPO/terraswap-service .
# run:
#   docker run --rm -it --env-file=path/to/.env --name terraswap-service -p 80:8080 REPO/terraswap-service

### BUILD
FROM golang:1.18-alpine AS build
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
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.a
RUN sha256sum /lib/libwasmvm_muslc.a | grep f6282df732a13dec836cda1f399dd874b1e3163504dbd9607c6af915b2740479

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

COPY --from=build /app/config.yaml /app/config.yaml
# Use an unprivileged user
USER appuser
COPY --from=build /app/cmd /app/cmd
# Grab compiled binary from build
COPY --from=build /app/main /app/main

# Expose application port
ENV APP_PORT 8000
EXPOSE $APP_PORT
# Set entry point
CMD [ "./main" ]
