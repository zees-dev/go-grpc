# BUILD
FROM golang:1.14 as build
LABEL maintainer="github.com/zees-dev"

# Install proto compiler and go protoc plugin
RUN apt update && \
  apt install -y protobuf-compiler && \
  go get github.com/golang/protobuf/protoc-gen-go

WORKDIR /go/src/app

COPY . .
RUN go mod download

# generate & build
RUN protoc --go_out=plugins=grpc:. ./api/protobuf-spec/hello.proto
RUN go build -o /go/bin/grpc

# MERGE
FROM gcr.io/distroless/base
COPY --from=build /go/bin/grpc /bin/
EXPOSE 50051
ENTRYPOINT [ "/bin/grpc" ]
CMD ["-type", "server"]