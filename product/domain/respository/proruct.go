package respository

import (
	"github.com/idcpj/micro/product/domain/model"
	"github.com/jinzhu/gorm"
	"log"
)

type IProductRepository interface {
	InitTable() error
	HasTable() bool
	FindProductByID(int64) (product *model.Product,err error)
	CreateProduct(product *model.Product) (int64,error)
	DeleteProductByID(id int64) error
	UpdateProduct(product *model.Product) error
	FindAll() ([]model.Product,error)
}
func NewProductRepository(db *gorm.DB) IProductRepository{
	return &ProductRepository{db: db}
}

type ProductRepository struct {
	db *gorm.DB
}

func (p *ProductRepository) InitTable() error {
	return p.db.CreateTable(&model.Product{},&model.ProductImage{},&model.ProductSize{},&model.ProductSeo{}).Error
}

func (p *ProductRepository) HasTable() bool {
	return p.db.HasTable(&model.Product{})
}

func (p *ProductRepository) FindProductByID(id int64) (*model.Product, error) {
	product := &model.Product{}
	err := p.db.Preload("ProductImage").
		Preload("ProductSize").Preload("ProductSeo").Where("id=?", id).Find(product).Error
	if err != nil {
		log.Println(err)
		return nil,err
	}
	return product,nil
}

func (p *ProductRepository) CreateProduct(product *model.Product) (int64, error) {
	tx:=p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println(r)
			return
		}
	}()

	if err :=p.db.Create(product).Error;err!=nil {
		tx.Rollback()
		log.Println(err)
		return 0,err
	}

	return  product.ID,tx.Commit().Error

}

func (p *ProductRepository) DeleteProductByID(id int64) error {
	tx := p.db.Begin()
	defer func() {
		if r:=recover();r!=nil {
			tx.Rollback()
		}
	}()

	if tx.Error==nil{
		return tx.Error
	}

	if err := tx.Unscoped().Where("id=?", id).Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("images_product_id=?", id).Delete(&model.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("size_product_id=?", id).Delete(&model.ProductSize{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("seo_product_id=?", id).Delete(&model.ProductSeo{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}

func (p *ProductRepository) UpdateProduct(product *model.Product) error {
	return p.db.Model(product).Update(product).Error
}

func (p *ProductRepository) FindAll() (products []model.Product,error error) {
	return products,p.db.Preload("ProductImage").
		Preload("ProductSize").Preload("ProductSeo").Find(&products).Error
}
