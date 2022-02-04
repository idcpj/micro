module github.com/idcpj/micro/product

go 1.15

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/asim/go-micro/plugins/config/source/consul/v4 v4.0.0-20220118152736-9e0be6c85d75
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.7.0
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20220118152736-9e0be6c85d75
	github.com/asim/go-micro/v3 v3.7.0
	github.com/jinzhu/gorm v1.9.16
	go-micro.dev/v4 v4.5.0
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	google.golang.org/protobuf v1.27.1
)
