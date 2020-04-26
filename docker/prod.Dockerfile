# BUILD
FROM golang:alpine as builder
LABEL maintainer="github.com/zees-dev"

ENV DEBIAN_FRONTEND=noninteractive

RUN apk add --update \
  git \
  protobuf
RUN go get github.com/golang/protobuf/protoc-gen-go

WORKDIR /go/src/grpc
COPY main.go go.mod go.sum ./
COPY internal internal
RUN protoc --go_out=plugins=grpc:. ./internal/protos/hello.proto
RUN GOARCH=386 GOOS=linux go build -ldflags="-w -s" -o /go/bin/grpc

# MERGE
FROM scratch
LABEL maintainer="github.com/zees-dev"

COPY --from=builder /go/bin/grpc /go/bin/grpc
EXPOSE 50051
ENTRYPOINT ["/go/bin/grpc","-type","server"]