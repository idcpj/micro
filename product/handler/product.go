package handler

import (
	"context"
	"github.com/idcpj/micro/common"
	"github.com/idcpj/micro/product/domain/model"
	"github.com/idcpj/micro/product/domain/service"
	product "github.com/idcpj/micro/product/proto"
	"log"
)

type Product struct{
	service service.IProductDataService
}

func NewProduct(service service.IProductDataService) *Product {
	return &Product{service: service}
}

func (p *Product) AddProduct(ctx context.Context, info *product.ProductInfo, product *product.ResponseProduct) error {
	pro:= &model.Product{}

	if err := common.SwapTo(info, pro);err != nil {
		log.Println(err)
		return err
	}

	id, err := p.service.AddProduct(pro)
	if err != nil {
		return err
	}
	product.Id=id
	return nil

}

func (p *Product) FindProductById(ctx context.Context, id *product.RequestId, info *product.ProductInfo) error {
	pro, err := p.service.FindProductById(id.GetProductId())
	if err != nil {
		return err
	}
	return common.SwapTo(pro,info)
}

func (p *Product) UpdateProduct(ctx context.Context, info *product.ProductInfo, response *product.Response) error {
	var pro *model.Product


	if err := common.SwapTo(info, pro);err != nil {
		return err
	}

	err := p.service.UpdateProduct(pro)
	if err != nil {
		return err
	}

	response.Msg="更新成功"
	return nil
}

func (p *Product) DeleteProductById(ctx context.Context, id *product.RequestId, response *product.Response) error {

	if err := p.service.DeleteProduct(id.GetProductId());err != nil {
		return err
	}
	response.Msg="删除成功"

	return nil
}

func (p *Product) FindAllProduct(ctx context.Context, all *product.RequestAll, product *product.AllProduct) error {

	pros, err := p.service.FindAllProduct()
	if err != nil {
		return err
	}
	return common.SwapTo(pros, &product.ProductInfo)
}

