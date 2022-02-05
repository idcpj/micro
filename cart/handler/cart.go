package handler

import (
	"context"
	"github.com/idcpj/micro/cart/common"
	"github.com/idcpj/micro/cart/domain/model"
	"github.com/idcpj/micro/cart/domain/service"
	pb "github.com/idcpj/micro/cart/proto/cart"
)

type Cart struct{
	Service service.ICartDataService
}

func (c *Cart) AddCart(ctx context.Context, req *pb.CartInfo, resp *pb.ResponseAdd)(err error) {

	cart :=&model.Cart{}
	common.SwapTo(req, cart)
	resp.CartId, err = c.Service.AddCart(cart)
	return err
}

func (c *Cart) CleanCart(ctx context.Context, req *pb.Clean, resp *pb.Response) error {
	cart:=&model.Cart{}
	common.SwapTo(req, cart)

	err := c.Service.ClearCart(req.UserId)
	if err != nil {
		return err
	}
	resp.Msg="购物车清空成功"

	return nil

}

func (c *Cart) Incr(ctx context.Context, req *pb.Item, resp *pb.Response) error {
	cart:=&model.Cart{}
	common.SwapTo(req, cart)
	err := c.Service.IncrNum(req.Id, req.ChangeNum)
	if err != nil {
		return err
	}
	resp.Msg="购物城数量增加成功"
	return nil

}

func (c *Cart) Decr(ctx context.Context, req *pb.Item, resp *pb.Response) error {
	cart:=&model.Cart{}
	common.SwapTo(req, cart)
	err:=c.Service.DecNum(req.Id,req.ChangeNum)
	if err != nil {
		return err
	}
	resp.Msg="购物城数量减少成功"
	return nil
}

func (c *Cart) DeleteItemById(ctx context.Context, req *pb.CartId, resp *pb.Response) error {
	cart:=&model.Cart{}
	common.SwapTo(req, cart)
	err := c.Service.DeleteCart(req.Id)
	if err != nil {
		return err
	}
	resp.Msg="删除成功"
	return nil
}

func (c *Cart) GetAll(ctx context.Context, req *pb.CartFindAll, resp *pb.CartAll) error {
	cart:=&model.Cart{}
	common.SwapTo(req, cart)
	allCart, err := c.Service.FindAllCart(req.UserId)
	if err != nil {
		return err
	}
	common.SwapTo(allCart,&resp.CartInfo)
	return nil
}

