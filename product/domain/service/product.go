package service

import (
	"github.com/idcpj/micro/product/domain/model"
	"github.com/idcpj/micro/product/domain/respository"
)

type IProductDataService interface {
	AddProduct(product *model.Product) (id int64,err error)
	DeleteProduct(id int64) error
	UpdateProduct(product *model.Product) error
	FindProductById(id int64) (*model.Product,error)
	FindAllProduct()([]model.Product,error)
}

func NewProductDataService(repository respository.IProductRepository) IProductDataService{
	return &ProductDataService{resp: repository}
}
type ProductDataService struct {
	resp  respository.IProductRepository
}

func (p *ProductDataService) AddProduct(product *model.Product) (id int64, err error) {
	return p.resp.CreateProduct(product)
}

func (p *ProductDataService) DeleteProduct(id int64) error {
	return p.resp.DeleteProductByID(id)
}

func (p *ProductDataService) UpdateProduct(product *model.Product) error {
	return p.resp.UpdateProduct(product)
}

func (p *ProductDataService) FindProductById(id int64) (*model.Product, error) {
	return p.resp.FindProductByID(id)
}

func (p *ProductDataService) FindAllProduct() ([]model.Product, error) {
	return p.resp.FindAll()
}

