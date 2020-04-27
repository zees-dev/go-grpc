# 1. Docker support

Date: 2020-27-04

## Status

Accepted

## Context

A [Dockerfile](../../Dockerfile) with accompanying [docker-compose.yml](../../docker-compose.yml) will be used to bootstrap the gRPC server and client.

## Decision

A docker-first approach is taken so that one can easily get the complete gRPC solution up and running with minimal effort - only docker and compose would be required.\
The provided [.dockerignore](../../.dockerignore) has been optimized to minimise docker context when building. A whitelist based approach is taken; this is very useful for large projects.\
The [Dockerfile](../../Dockerfile) adheres to best practices Docker practices via making use of [multi-stage](https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds) docker builds and secure [distroless docker images](https://github.com/GoogleContainerTools/distroless).

## Consequences

Docker and docker-compose are required to be installed on the user machine.
