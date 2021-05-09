package main

import (
	"category/proto/category"
	"context"
	"github.com/asim/go-micro/v3"
)

func main() {

	service := micro.NewService(
		micro.Name("category.client"),
		micro.Version("latest"),
	)

	service.Init()

	categoryService := category.NewCategoryService("category.client", service.Client())
	categoryService.CreateCategory(context.Background(),)

}