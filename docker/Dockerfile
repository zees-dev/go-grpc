# BUILD
FROM golang:1.14 as build
LABEL maintainer="github.com/zees-dev"

RUN apt update && \
  apt install -y protobuf-compiler
RUN go get github.com/golang/protobuf/protoc-gen-go

WORKDIR /go/src/app

# cache depedency downloads
COPY main.go go.mod go.sum ./
RUN go mod download
COPY internal internal

# generate & build
RUN protoc --go_out=plugins=grpc:. ./internal/protos/hello.proto
RUN go build -o /go/bin/app

# MERGE
FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /bin/
EXPOSE 50051
ENTRYPOINT [ "/bin/app" ]
CMD ["-type", "server"]