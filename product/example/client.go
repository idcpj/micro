package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/idcpj/micro/product/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"log"
)

func main() {

	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})

	srv := micro.NewService(
		micro.Name("product.client"),
		micro.Version("latest"),
		micro.Registry(consulReg),
	)
	srv.Init()

	service := proto.NewProductService("product.service", srv.Client())

	//addReq := &proto.ProductInfo{}
	//addReq.ProductName = "name:" + strconv.Itoa(rand.Int())
	//addReq.ProductDescription = "this is a desc"
	//addReq.ProductPrice = 123.123
	//addReq.ProductSku = "sku sku:" + strconv.Itoa(rand.Int())
	//addReq.ProductSeo = &proto.ProductSeo{
	//	SeoTitle:       "title",
	//	SeoKeywords:    "keyword",
	//	SeoDescription: "desc",
	//	SeoCode:        "code"+ strconv.Itoa(rand.Int()),
	//}
	//
	//var size []*proto.ProductSize
	//size = append(size, &proto.ProductSize{
	//	SizeName: "size_name",
	//	SizeCode: "code 1"+ strconv.Itoa(rand.Int()),
	//})
	//size = append(size, &proto.ProductSize{
	//	SizeName: "size_name11",
	//	SizeCode: "code 122"+ strconv.Itoa(rand.Int()),
	//})
	//addReq.ProductSize = size
	//
	//product, err := service.AddProduct(context.Background(), addReq)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Printf("%+v\n", product.Id)

	findReq := &proto.RequestId{ProductId: 3}

	info, err := service.FindProductById(context.Background(), findReq)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", info.ProductName)


}
