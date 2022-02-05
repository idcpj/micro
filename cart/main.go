package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/idcpj/micro/cart/domain/repository"
	service2 "github.com/idcpj/micro/cart/domain/service"
	"github.com/idcpj/micro/cart/handler"
	"github.com/idcpj/micro/cart/proto/cart"
	"github.com/idcpj/micro/common"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var (
	QPS=100
)

func main() {
	// 配置中心
	configConsul, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}



	// 注册 ccart  成服务需要的consul 服务地址
	consulRegistry:=consul.NewRegistry(func(opts *registry.Options) {
		opts.Addrs=[]string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪
	tracer, closer, err := common.NewTracer(
		"go.micro.service.cart",
		"127.0.0.1:6831", // jaeger 的地址
	)
	if err != nil {
		log.Error(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// 数据库
	mysqlInfo := common.GetMysqlFromConsul(configConsul, "mysql")

	//db, err := gorm.Open("mysql", "root:12345678@/micro?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlInfo.User,mysqlInfo.Pwd,mysqlInfo.Database))
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	db.SingularTable(true)
	db.LogMode(true)

	// 初始化 repository
	cartRepository := repository.NewCartRepository(db)
	err = cartRepository.InitTable()
	if err != nil {
		log.Error(err)
	}

	// Create service
	srv := micro.NewService(
		micro.Name("cart.server"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:807"),
		micro.Registry(consulRegistry),

		micro.WrapHandler(
			// 链路追踪
			opentracing2.NewHandlerWrapper(opentracing.GlobalTracer()),

			// 添加限流
			ratelimit.NewHandlerWrapper(QPS),
		),

	)
	srv.Init()

	cartServer := service2.NewCartDataServer(cartRepository)


	// Register handler
	cart.RegisterCartHandler(srv.Server(), &handler.Cart{Service: cartServer})

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
