GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/category.proto
.PHONY: build
build:
	go build -o category *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY:consul
consul:
	docker run -d --name consul  -p 8500:8500  consul

.PHONY: docker
docker:
	docker build . -t category:latest
