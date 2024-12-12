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
ADD https://github.com/CosmWasm/wasmvm/releases/download/v2.1.4/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v2.1.4/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
#ADD https://github.com/CosmWasm/wasmvm/releases/download/v2.1.4/libwasmvmstatic_darwin.a /lib/libwasmvm_muslc.darwin.a
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 090b97641157fae1ae45e7ed368a1a8c091f3fef67958d3bc7c2fa7e7c54b6b4
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep a4a3d09b36fabb65b119d5ba23442c23694401fcbee4451fe6b7e22e325a4bac
#RUN sha256sum /lib/libwasmvm_muslc.darwin.a | grep f7361db078218eb31ebaeab71a6c242034a3709886848596078d132ac7e08e36
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