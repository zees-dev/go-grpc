# Go gRPC

A simple gRPC client and server written in go.\
This project is meant to be a template to allow one to quickly bootstrap a gRPC based project.

## Build

- Generate Protobuf source files:

  ```sh
  protoc --go_out=plugins=grpc:. ./internal/protos/hello.proto
  ```

- Build solution:

  ```sh
  go build main.go
  ```

- Run Server:

  ```sh
  ./main -type server
  ```

- Run client - with message _"my-message"_:

  ```sh
  ./main -type client my-message
  ```

## Docker

Run server and client via docker-compose:

  ```sh
  docker-compose up --build
  ```

Note: The client continously exits when gRPC server successfully recieves message from client

## TODO

- [x] Server implementation
- [x] Client implementation
- [ ] Unit Tests
- [x] Containerisation (Docker)
- [ ] Github actions (CICD)
- [ ] Docs

### Bugs

- [ ] client prints out address flags when passed in

## Tests

  ```sh
  go test
  ```

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
