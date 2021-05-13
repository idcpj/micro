package main

import (
	"github.com/idcpj/micro/category/proto/category"
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	service := micro.NewService(
		micro.Name("category.client"),
		micro.Version("latest"),
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
	levelResp, err := categoryService.FindCategoryByLevel(context.Background(), req3, client.WithRequestTimeout(30*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("len levelResp %+v\n", len(levelResp.Category))

}
