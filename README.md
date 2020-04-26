# Go gRPC

A simple gRPC client and server written in go.\
This project is meant to be a template to allow one to quickly bootstrap a gRPC based project.

## Build

1. Generate Protobuf source files:\
  ```protoc --go_out=plugins=grpc:. ./internal/protos/hello.proto```

2. Build & Run solution:\
  ```go run main.go```

## TODO

- [x] Server implementation
- [x] Client implementation
- [ ] Unit Tests
- [ ] Containerisation (Docker)
- [ ] Github actions (CICD)
- [ ] Docs

## Tests

```go test```

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
