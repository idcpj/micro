package repository

import (
	"errors"
	"github.com/idcpj/micro/cart/domain/model"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartById(int64) (*model.Cart,error)
	CreatCart(cart *model.Cart)(int64,error)
	DeleteCartById(int64) error
	UpdateCart(cart *model.Cart)error
	FindAll(userId int64)([]*model.Cart,error)

	ClearCart(int64) error
	IncrNum(int64,int64) error
	DecNum(int64,int64) error

}

func NewCartRepository(db *gorm.DB)ICartRepository{
	return &CartRepository{
		mysqldb: db,
	}
}

type CartRepository struct {
	mysqldb *gorm.DB
}


func (c *CartRepository) InitTable() error {
	return c.mysqldb.CreateTable(&model.Cart{}).Error
}
func (c *CartRepository) ClearCart(userId int64) error {
	return c.mysqldb.Where("user_id=?",userId).Delete(&model.Cart{}).Error
}

func (c *CartRepository) IncrNum(cartId int64, num int64) error {
	cart:=&model.Cart{
		ID:        cartId,
	}
	return c.mysqldb.Model(cart).UpdateColumn("num",
		gorm.Expr("num + ?",num)).Error
}

func (c *CartRepository) DecNum(cartId int64, num int64) error {
	cart :=&model.Cart{
		ID: cartId,
	}
	db := c.mysqldb.Model(cart).Where("num>=?", num).UpdateColumn("num", gorm.Expr("num  - ?", num))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}
	return nil
}


func (c *CartRepository) FindCartById(carId int64) (*model.Cart, error) {
	cart := &model.Cart{}
	return cart,c.mysqldb.First(cart,carId).Error
}

func (c *CartRepository) CreatCart(cart *model.Cart) (int64, error) {
	db:=c.mysqldb.FirstOrCreate(cart,model.Cart{
		ProductId: cart.ProductId,
		SizeID:    cart.SizeID,
		UserID: cart.UserID,
	})
	if db.Error != nil {
		return 0,db.Error
	}
	if db.RowsAffected == 0 {
		return 0,errors.New("购物车插入失败")
	}
	return cart.ID,nil
}

func (c *CartRepository) DeleteCartById(cartId int64) error {
	return c.mysqldb.Where("id=?",cartId).Delete(&model.Cart{}).Error
}

func (c *CartRepository) UpdateCart(cart *model.Cart) error {
	return c.mysqldb.Model(cart).Update(cart).Error
}

func (c *CartRepository) FindAll(userId int64) (carts []*model.Cart,err error) {
	return carts,c.mysqldb.Where("user_id=?",userId).Find(&carts).Error
}
