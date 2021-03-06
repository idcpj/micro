package main

import (
	"github.com/idcpj/micro/user/domain/repository"
	"github.com/idcpj/micro/user/domain/service"
	"github.com/idcpj/micro/user/handler"
	"github.com/idcpj/micro/user/proto/user"
	"github.com/jinzhu/gorm"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("user.server"),
		micro.Version("latest"),
	)
	srv.Init()

	db, err := gorm.Open("mysql", "root:12345678@/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//db.LogMode(true)


	// gorm会在创建表的时候去掉”s“的后缀
	db.SingularTable(true)


	userRepository := repository.NewUserRepository(db)

	if !userRepository.HasTable() {
		// 只执行一次
		if err := userRepository.InitTable(); err != nil {
			log.Fatal(err)
		}
	}

	dataService := service.NewUserDataService(userRepository)


	// Register handler
	if err := user.RegisterUserHandler(srv.Server(), handler.NewUser(dataService)); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
