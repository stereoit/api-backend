#!make
include .env
export $(shell sed 's/=.*//' .env)

REST_BINARY=rest-server
GRPC_BINARY=grpc-server

.PHONY: grpc rest build test clean docker run-docker grpcc protoc mocks

grpc:
	go run cmd/grpcserver/main.go

rest:
	go run cmd/rest/main.go

build:
	go build -o ${REST_BINARY} cmd/rest/main.go
	go build -o ${GRPC_BINARY} cmd/grpcserver/main.go

clean:
	@rm -f ${REST_BINARY} 
	@rm -f ${GRPC_BINARY}

test:
	go test -cover -covermode=atomic ./...

integration-test:
	go test -count=1 -tags=integration ./...

docker:
	podman build -t eventival-backend .

run-docker:
	podman run -it  -e REST_PORT=8000 -e STORAGE_BACKEND=MEMORY --rm -p 8000:8000 eventival-backend

grpcc:
	grpcc --proto pkg/user/interface/rpc/v1.0/protocol/*.proto --address 127.0.0.1:${GRPC_PORT} -i

protoc:
	protoc --proto_path=. --go_out=plugins=grpc:./ pkg/user/interface/rpc/v1.0/protocol/*.proto

mocks: 
	mockery -dir pkg/user/domain -recursive -all -output pkg/mocks