package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"go-micro"
	"github.com/idcpj/micro/user/proto/user"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags|log.Lshortfile)
}
func main() {


	service := micro.NewService(
		micro.Name("user.client"),
		micro.Version("latest"),
	)
	service.Init()

	userService := user.NewUserService("user.server", service.Client())
	req:=&user.UserRegisterRequest{}
	req.FirstName="chen2"+strconv.Itoa(rand.Int())
	req.UserName="pengjie3"+strconv.Itoa(rand.Int())
	req.Pwd="123"

	registerResp, err := userService.Register(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("regiester response %+v\n", registerResp.GetMessage())

	infoReq:=&user.UserInfoRequest{}
	infoReq.UserName="pengjie"
	infoResp, err := userService.GetUserInfo(context.TODO(), infoReq,client.WithRequestTimeout(30*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("GetUserInfo:%+v\n", infoResp)

}