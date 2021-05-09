package main

import (
	"category/domain/repository"
	"category/domain/service"
	"category/handler"
	"category/proto/category"
	"github.com/asim/go-micro/v3"
	"github.com/jinzhu/gorm"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("category.service"),
		micro.Version("latest"),
	)
	srv.Init()

	db, err := gorm.Open("mysql", "root:12345678@/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.SingularTable(true)
	db.LogMode(true)

	respository := repository.NewCategoryRespository(db)
	if !respository.HasTable() {
		if err := respository.InitTable(); err != nil {
			log.Fatalln(err)
		}
	}
	dataService := service.NewCateGoryDataService(respository)
	cateHandle := handler.NewCategory(dataService)
	// Register handler

	if err := category.RegisterCategoryHandler(srv.Server(), cateHandle); err != nil {
		log.Fatalln(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
