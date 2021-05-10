module category

go 1.15

require (
	github.com/asim/go-micro/cmd/protoc-gen-micro/v3 v3.0.0-20210506054305-9e9157d878dc // indirect
	github.com/asim/go-micro/v3 v3.5.1
	github.com/golang/protobuf v1.5.2
	github.com/jinzhu/gorm v1.9.16
	google.golang.org/protobuf v1.26.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
