
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/asim/go-micro/v3/cmd/protoc-gen-micro
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/user.proto
	# mac
	 sed  -i ""  's|micro/micro/v3/service|asim/go-micro/v3|g' proto/**/*.micro.go
	# linux
	#sed  -i   's|micro/micro/v3/service|asim/go-micro/v3|g' proto/**/*.micro.go
.PHONY: build
build:
	go build -o user *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user:latest
