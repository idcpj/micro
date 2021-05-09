package main

import (
	"category/handler"
	"category/proto/category"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("category"),
		service.Version("latest"),
	)

	// Register handler
	category.RegisterCategoryHandler(srv.Server(), new(handler.Category))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
