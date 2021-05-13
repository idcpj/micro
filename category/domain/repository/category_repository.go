package repository

import (
	"github.com/idcpj/micro/category/domain/model"
	"github.com/jinzhu/gorm"
)

type ICateGoryRepository interface {
	InitTable() error
	HasTable() bool
	FindCategoryById(int64) (*model.Category, error)
	CreateCategory(*model.Category) (int64, error)
	DeleteCategoryById(int64) error
	UpdateCategory(*model.Category) error
	FindALl() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

func NewCategoryRespository(db *gorm.DB) ICateGoryRepository {
	return &CategoryRespository{
		db: db,
	}
}

type CategoryRespository struct {
	db *gorm.DB
}

func (c *CategoryRespository) HasTable() bool {
	return c.db.HasTable(new(model.Category))
}

func (c *CategoryRespository) FindCategoryByName(name string) (category *model.Category, err error) {
	return category, c.db.Where("category_name=?", name).Find(category).Error
}

func (c *CategoryRespository) FindCategoryByLevel(levelId uint32) ([]model.Category, error) {
	var category []model.Category
	return category, c.db.Where("category_level=?", levelId).Find(&category).Error
}

func (c *CategoryRespository) FindCategoryByParent(parentId int64) ([]model.Category, error) {
	var category []model.Category
	return category, c.db.Where("category_parent=?", parentId).Find(category).Error
}

func (c *CategoryRespository) InitTable() error {
	return c.db.CreateTable(&model.Category{}).Error
}

func (c *CategoryRespository) FindCategoryById(id int64) (*model.Category, error) {
	cate := &model.Category{}
	return cate, c.db.Where("id=?", id).Find(cate).Error
}

func (c *CategoryRespository) CreateCategory(category *model.Category) (int64, error) {
	return category.ID, c.db.Create(category).Error
}

func (c *CategoryRespository) DeleteCategoryById(id int64) error {
	return c.db.Where("id=?", id).Delete(&model.Category{}).Error
}

func (c *CategoryRespository) UpdateCategory(category *model.Category) error {
	return c.db.Model(category).Update(category).Error
}

func (c *CategoryRespository) FindALl() (cates []model.Category, err error) {
	return cates, c.db.Find(cates).Error
}
