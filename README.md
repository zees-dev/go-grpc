# Go gRPC

A simple gRPC client and server written in go.\
Both the client and server use the same code generated by the protobuf definition (proto3).

This project is meant to be a template to allow one to quickly bootstrap a gRPC based project.

All architectural decisions are documented as architecture design records (ADR's) in the **docs/adr** directory.

Note: This project is a WIP and will continue to grow as expertise in gRPC based microservice architecture develops.

---

## Local Setup

### Pre-requisites

- [protoc compiler](https://github.com/protocolbuffers/protobuf/releases) - pre-compiled binary for your OS
  - Note: add this to your path
- **protoc-gen-go** - protoc plugin for go
  - ```go get github.com/golang/protobuf/protoc-gen-go```

### Build

- Generate Protobuf source files:

  ```sh
  protoc --go_out=plugins=grpc:. ./api/protobuf-spec/hello.proto
  ```

- Build solution:

  ```sh
  go build -o grpc main.go
  ```

### Run

- Run Server:

  ```sh
  ./grpc -type server
  ```

- Run client - with message _"my-message"_:

  ```sh
  ./grpc -type client my-message
  ```

---

## Docker

Run server and client via docker-compose:

  ```sh
  docker-compose up --build
  ```

Note: The client continously exits when gRPC server successfully recieves message from client

---

## TODO

- [x] Server implementation
- [x] Client implementation
- [ ] Unit Tests
- [x] Containerisation (Docker)
- [x] Create/Use distroless based prod docker image
- [x] Docs - Architecture Design Records
- [x] Godoc support
- [x] Setup Makefile
- [ ] Github actions (CICD)

---

## Tests

  ```sh
  go test
  ```

---

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
