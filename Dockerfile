ARG IMG_TAG=latest
ARG PLATFORM="linux/amd64"
ARG GO_VERSION="1.23.1"
ARG RUNNER_IMAGE="gcr.io/distroless/static"

FROM --platform=${PLATFORM} golang:${GO_VERSION}-alpine3.20 as builder
WORKDIR /src/app/
COPY go.mod go.sum* ./
RUN go mod download
COPY . .

# From https://github.com/CosmWasm/wasmd/blob/master/Dockerfile
# For more details see https://github.com/CosmWasm/wasmvm#builds-of-libwasmvm
ARG ARCH=x86_64
# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v2.1.3/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v2.1.3/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
#ADD https://github.com/CosmWasm/wasmvm/releases/download/v2.1.3/libwasmvmstatic_darwin.a /lib/libwasmvm_muslc.darwin.a
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep faea4e15390e046d2ca8441c21a88dba56f9a0363f92c5d94015df0ac6da1f2d
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep 8dab08434a5fe57a6fbbcb8041794bc3c31846d31f8ff5fb353ee74e0fcd3093
#RUN sha256sum /lib/libwasmvm_muslc.darwin.a | grep f7a997c6a769e5624dac910dc2f0bec4c386aeb54342dd04ef6d5eba0340a20d
RUN cp /lib/libwasmvm_muslc.${ARCH}.a /lib/libwasmvm_muslc.a

ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES
RUN set -eux; apk add --no-cache ca-certificates build-base;

ARG VERSION=""
RUN BUILD_TAGS=muslc LINK_STATICALLY=true LDFLAGS=-buildid=$VERSION make build

# Add to a distroless container
ARG PLATFORM="linux/amd64"
FROM --platform=${PLATFORM} gcr.io/distroless/cc:$IMG_TAG
ARG IMG_TAG
COPY --from=builder /src/app/build/bcnad /usr/local/bin/bcnad

EXPOSE 26656 26657 1317 9090

ENTRYPOINT ["bcnad"]