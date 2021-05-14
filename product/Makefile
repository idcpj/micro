GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/product.proto
	
.PHONY: build
build:
	go build -o product *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t product:latest
