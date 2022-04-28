package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/idcpj/micro/cart/proto/cart"
	"strconv"
)

type CartApi struct {
	CatService cart.CartService
}

func (c *CartApi) FindAll(gin *gin.Context) {

	userid_s, _ := gin.GetQuery("userid")
	userid, err := strconv.Atoi(userid_s)
	if err != nil {
		gin.String(501, err.Error(), "aaa")
	}
	cartAll := &cart.CartFindAll{
		UserId: int64(userid),
	}

	all, err := c.CatService.GetAll(context.Background(), cartAll)
	if err != nil {
		gin.String(501, err.Error(), "aaa")
	}
	gin.JSON(200, all.CartInfo)

}
