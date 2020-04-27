# 1. gRPC services

Date: 2020-27-04

## Status

Accepted

## Context

Create gRPC server and client in the same repo so the project can be used as a template for future, larger projects

## Decision

Use Go to develop a gRPC (with Protocol buffers) server and client. We will use the [golang protoc plugin - **protoc-gen-go**](https://github.com/golang/protobuf/tree/master/protoc-gen-go) to generate Go gRPC service code.

## Consequences

- This project assumes both Server and Client as coded in Go - however code can be generated for other languages too
  - See supported languages [here](https://github.com/protocolbuffers/protobuf)
