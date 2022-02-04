module github.com/idcpj/micro/category

go 1.15

require (
	github.com/asim/go-micro/plugins/config/source/consul/v4 v4.0.0-20220118152736-9e0be6c85d75
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.7.0
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20220118152736-9e0be6c85d75
	github.com/jinzhu/gorm v1.9.16
	go-micro.dev/v4 v4.2.1
	google.golang.org/protobuf v1.26.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
