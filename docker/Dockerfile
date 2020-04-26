# BUILD
FROM golang:alpine as builder
LABEL maintainer="github.com/zees-dev"

ENV DEBIAN_FRONTEND=noninteractive

RUN apk add --update \
  git \
  protobuf \
  && go get github.com/golang/protobuf/protoc-gen-go

WORKDIR /go/src/grpc
COPY main.go go.mod go.sum ./
COPY internal internal
RUN protoc --go_out=plugins=grpc:. ./internal/protos/hello.proto && \
  go build main.go && \
  chmod +x ./main

# Generate entrypoint
RUN echo '#!/bin/sh' > /docker-entrypoint.sh && \
  echo 'set -ex' >> /docker-entrypoint.sh && \
  echo 'exec "$@"' >> /docker-entrypoint.sh && \
  chmod +x /docker-entrypoint.sh

EXPOSE 50051
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["/go/src/grpc/main","-type","server"]