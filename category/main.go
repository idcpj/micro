package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/idcpj/micro/category/common"
	"github.com/idcpj/micro/category/domain/repository"
	"github.com/idcpj/micro/category/domain/service"
	"github.com/idcpj/micro/category/handler"
	"github.com/idcpj/micro/category/proto/category"
	"github.com/jinzhu/gorm"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// 配置中心
	// 8500 为 consul的服务地址,先链接 consul 的服务
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Fatalln(err)
	}


	// 注册 consul 成服务需要的consul 服务地址
	consulRegistry:=consul.NewRegistry(func(opts *registry.Options) {
		opts.Addrs=[]string{
			"127.0.0.1:8500",
		}
	})

	// Create service
	srv := micro.NewService(
		micro.Name("category.service"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8082"), // 指定此服务的端口
		micro.Registry(consulRegistry), // 把 micro 注册到 consul
		micro.Config(consulConfig),
	)
	srv.Init()


	// 读取配置的两种方式
	// 方式一:直接从consul获取mysql 配置,从 consul 中获取mysql的配置信息
	//mysqlInfo :=common.GetMysqlFromConsul(consulConfig,"mysql")

	// 方式二: 从 注册的微服务中获取,前提是必须注册成 micro.Config(consulConfig),
	mysqlInfo :=common.GetMysqlFromConsul(srv.Options().Config,"mysql")

	//db, err := gorm.Open("mysql", "root:12345678@/micro?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",mysqlInfo.User,mysqlInfo.Pwd,mysqlInfo.Database))
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
