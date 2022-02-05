package service

import (
	"github.com/idcpj/micro/cart/domain/model"
	"github.com/idcpj/micro/cart/domain/repository"
)

type ICartDataService interface {
	AddCart(cart *model.Cart) (int64,error)
	DeleteCart(int64)error
	UpdateCart(cart *model.Cart)error
	FindCartById(int64) (*model.Cart,error)
	FindAllCart(userId int64)([]*model.Cart,error)

	ClearCart(int64) error
	IncrNum(int64,int64) error
	DecNum(int64,int64) error
}

type CartDataServer struct {
	repo repository.ICartRepository
}

func NewCartDataServer(repo repository.ICartRepository) ICartDataService {
	return &CartDataServer{repo: repo}
}

func (c *CartDataServer) AddCart(cart *model.Cart) (int64, error) {
	return c.repo.CreatCart(cart)
}

func (c *CartDataServer) DeleteCart(cartid int64) error {
	return c.repo.DeleteCartById(cartid)
}

func (c *CartDataServer) UpdateCart(cart *model.Cart) error {
	return c.repo.UpdateCart(cart)
}

func (c *CartDataServer) FindCartById(cartId int64) (*model.Cart, error) {
	return c.repo.FindCartById(cartId)
}

func (c *CartDataServer) FindAllCart(userId int64) ([]*model.Cart, error) {
	return c.repo.FindAll(userId)
}

func (c *CartDataServer) ClearCart(userId int64) error {
	return c.repo.ClearCart(userId)
}

func (c *CartDataServer) IncrNum(cartId int64, num int64) error {
	return c.repo.IncrNum(cartId,num)
}

func (c *CartDataServer) DecNum(cartId int64, num int64) error {
	return c.repo.DecNum(cartId,num)
}

