package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/idcpj/micro/cart/proto/cart"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {

	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8700"}
	})

	srv := micro.NewService(
		micro.Name("product.client"),
		micro.Version("latest"),
		micro.Registry(consulReg),
	)
	srv.Init()

	service := cart.NewCartService("cart.server", srv.Client())
	req := &cart.CartInfo{
		Id:        1231,
		UserId:    11,
		ProductId: 222,
		SizeId:    222,
		Num:       3333,
	}
	addCart, err := service.AddCart(context.TODO(), req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", addCart)

}
