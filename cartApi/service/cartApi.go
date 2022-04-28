package service

//
//import (
//	pb "cartApi/proto/cartApi"
//	"context"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/idcpj/micro/cart/proto/cart"
//	log "go-micro.dev/v4/logger"
//	"strconv"
//)
//
//type ICartApi interface {
//	FindAll (g *gin.Context)
//	//CleanCart (g *gin.Context)
//	//Incr (g *gin.Context)
//	//Decr (g *gin.Context)
//	//DeleteItemById (g *gin.Context)
//	//GetAll (g *gin.Context)
//}
//
//var CartApi cartApi
//
//type cartApi struct{
//	CatService cart.CartService
//}
//
//
//func (c *cartApi) FindAll(g *gin.Context)  {
//	log.Info("接受到访问请求")
//	if _, ok := req.Get["user_id"];!ok {
//		resp.StatusCode=500
//		return errors.New("参数异常")
//	}
//	userIdString :=req.Get["user_id"].Values[0]
//	fmt.Printf("%+v\n", userIdString)
//
//	userId, err := strconv.Atoi(userIdString)
//	if err != nil {
//		return err
//	}
//	// 获取给购物车所有商品
//	cartAll, err := c.CatService.GetAll(context.TODO(),
//		&cart.CartFindAll{UserId: int64(userId)},
//	)
//	if err != nil {
//		return err
//	}
//	b, err := json.Marshal(cartAll)
//	if err != nil {
//		return err
//	}
//	resp.StatusCode=200
//	resp.Body=string(b)
//	return nil
//}
//
