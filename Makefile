#!make
include .env
export $(shell sed 's/=.*//' .env)

run:
	go run cmd/grpcserver/main.go

rest:
	go run cmd/rest/main.go

grpcc:
	grpcc --proto pkg/user/interface/rpc/v1.0/protocol/*.proto --address 127.0.0.1:${GRPC_PORT} -i

protoc:
	protoc --proto_path=. --go_out=plugins=grpc:./ pkg/user/interface/rpc/v1.0/protocol/*.proto
