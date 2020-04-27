all: codegen test clean build

codegen:
	protoc --go_out=plugins=grpc:. ./internal/protos/hello.proto

build:
	protoc --go_out=plugins=grpc:. ./internal/protos/hello.proto
	go build -o grpc main.go

clean:
	rm -rf grpc
	rm -rf internal/protos/**.pb.go

test:
	# go test -v -short ./...
	echo "TODO implement testing..."
.PHONY: test

server:
	go run main.go -type server

client:
	go run main.go -type client

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down

doc:
	godoc -http localhost:6060
.PHONY: doc
