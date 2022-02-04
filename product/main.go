package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/idcpj/micro/product/common"
	"github.com/idcpj/micro/product/domain/respository"
	"github.com/idcpj/micro/product/domain/service"
	"github.com/idcpj/micro/product/handler"
	product "github.com/idcpj/micro/product/proto"
	"github.com/jinzhu/gorm"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
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
		micro.Name("product.service"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8083"), // 指定此服务的端口
		micro.Registry(consulRegistry), // 把 micro 注册到 consul
		micro.Config(consulConfig),
	)

	srv.Init()

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

	resp := respository.NewProductRepository(db)
	if !resp.HasTable() {
		if err := resp.InitTable(); err != nil {
			log.Fatalln(err)
		}
	}

	dataService := service.NewProductDataService(resp)

	// Register handler
	if err := product.RegisterProductHandler(srv.Server(), handler.NewProduct(dataService));err!=nil{
		log.Fatalln(err)
	}


	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
