# syntax = docker/dockerfile:1-experimental
FROM --platform=linux/amd64 golang:1.15.3-alpine AS base

WORKDIR /src

# Download required modules to allow Docker to cache these modules as it 
# will only rerun these steps if the go.* files change.
COPY go.* .
RUN go mod download

FROM base as build
    # Bind mounting from context saves time compared to COPY 
RUN --mount=target=. \ 
    #
    --mount=type=cache,target=/root/.cache/go-build \
    # Some Go packages might use CGO under the hood, leveraging common C 
    # libraries for broader compatibility. These libraries are found on 
    # most OS distributions - but obviously not in the Scratch image. CGO 
    # is deliberately disabled to ensure that the resulting binary will not 
    # be linked to any C libraries.
    CGO_ENABLED=0 \
    go build -o /go/bin/kimbap-server ./cmd/kimbap-server

FROM golangci/golangci-lint:v1.27-alpine AS lint-base

# Go linter runner target
FROM base AS lint
RUN --mount=target=. \
    --mount=from=lint-base,src=/usr/bin/golangci-lint,target=/usr/bin/golangci-lint \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/.cache/golangci-lint \
    CGO_ENABLED=0 \
    golangci-lint run --timeout 10m0s ./...

# Unit tests target
FROM base AS unit-tests
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 \
    go test -v ./...

# Copy binary to Docker minimalistic scratch image 
FROM scratch as bin
COPY --from=build /go/bin/kimbap-server /go/bin/kimbap-server

ENTRYPOINT [ "/go/bin/kimbap-server" ]
