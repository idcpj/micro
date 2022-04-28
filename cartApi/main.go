package main

import (
	"cartApi/handler"
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	hystrix2 "github.com/asim/go-micro/plugins/wrapper/breaker/hystrix/v4"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v4"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/gin-gonic/gin"
	"github.com/idcpj/micro/cart/proto/cart"
	"github.com/idcpj/micro/common"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"net"
	"net/http"
)

var (
	service = "cart.api"
	version = "latest"
	addr    = "0.0.0.0:8096"
)

func main() {
	// 配置中心
	configConsul, err := common.GetConsulConfig("127.0.0.1", 8700, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	defer configConsul.Close()

	// 链路追踪
	tracer, closer, err := common.NewTracer("go.micro.api.cartApi", "127.0.0.1:831")
	if err != nil {
		log.Error(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	//熔断器
	hystrixHandler := hystrix.NewStreamHandler()
	hystrixHandler.Start()
	// 启动端口
	go func() {
		err := http.ListenAndServe(net.JoinHostPort("", "9096"), hystrixHandler)
		if err != nil {
			panic(err)
		}
	}()

	// 注册 cart  成服务需要的consul 服务地址
	consulRegistry := consul.NewRegistry(func(opts *registry.Options) {
		opts.Addrs = []string{
			"127.0.0.1:8700",
		}
	})

	// 同时控制退出
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create service
	srv := micro.NewService(
		micro.Context(ctx),

		micro.Config(configConsul),
		micro.Registry(consulRegistry),

		//web.RegisterTTL(30 * time.Second),//设置注册服务的过期时间
		//web.RegisterInterval(20 * time.Second),//设置间隔多久再次注册服务

		// 链路追踪
		// 因为是以客户端的形式访问,
		// WrapHandler -> WrapClient
		micro.WrapClient(
			opentracing2.NewClientWrapper(opentracing.GlobalTracer()),
		),

		// 添加熔断 micro v4 可以直接添加
		// 被调用端设置一个限流值, 调用端通过熔断机制,去判断是否熔断
		micro.WrapClient(hystrix2.NewClientWrapper()),

		// 添加负责均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)
	srv.Init()

	// router
	g := gin.New()
	cartService := cart.NewCartService("cart.server", srv.Client())
	handlers := &handler.CartApi{
		CatService: cartService,
	}
	g.GET("/cartApi/Findall", handlers.FindAll)

	// web micro
	newService := web.NewService(
		web.Name(service),
		web.Version(version),
		web.Context(ctx),
		web.Address(addr),
		web.MicroService(srv),
		web.HandleSignal(true),
		web.Handler(g),
	)

	err = newService.Init()
	if err != nil {
		log.Error(err)
	}

	err = newService.Run()
	if err != nil {
		log.Errorf("web run error: %v", err)
	}

}
