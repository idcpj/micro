module github.com/idcpj/micro/product

go 1.15

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/asim/go-micro/plugins/config/source/consul/v3 v3.0.0-20210523073820-acc3f5479f48
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.0.0-20210523073820-acc3f5479f48
	github.com/asim/go-micro/v3 v3.5.1
	github.com/golang/protobuf v1.5.2
	github.com/jinzhu/gorm v1.9.16
	google.golang.org/protobuf v1.26.0
)