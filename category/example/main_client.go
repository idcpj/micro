package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/idcpj/micro/category/proto/category"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"log"
	"math/rand"
	"strconv"
)

func main() {


	// 配置注册中心就可以访问配置中心
	consulRegistry:=consul.NewRegistry(func(opts *registry.Options) {
		opts.Addrs=[]string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Name("category.client"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)

	service.Init()

	categoryService := category.NewCategoryService("category.service", service.Client())
	req1 := &category.CategoryRequest{}
	req1.CategoryLevel = 1
	req1.CategoryImage = "/imgr/1"
	req1.CategoryDescription = "this is a desc"
	req1.CategoryParent = 0
	req1.CategoryName = "test" + strconv.Itoa(rand.Int())

	// create
	createResp, err := categoryService.CreateCategory(context.Background(), req1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", createResp)

	req2 := &category.FindByIdRequest{}
	req2.Id = createResp.CategoryId

	// find by id
	IdResp, err := categoryService.FindCategoryByID(context.Background(), req2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", IdResp)

	// find  by level_id
	req3 := &category.FindByLevelRequest{}
	req3.CategoryLevel = 1
	levelResp, err := categoryService.FindCategoryByLevel(context.Background(), req3)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("len levelResp %+v\n", len(levelResp.Category))

}
