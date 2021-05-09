package service

import (
	"category/domain/model"
	"category/domain/repository"
)

type ICategoryDataService interface {
	AddCategory(model *model.Category) (int64,error)
	DeleteCategory(id int64) error
	UpdateCategory(*model.Category) error
	FindCategoryById(int64) (*model.Category,error)
	FindAllCategory()([]*model.Category,error)
	FindCategoryByName(string)(*model.Category,error)
	FindCategoryByLevel(uint32)([]*model.Category,error)
	FindCategoryByParent(int64)([]*model.Category,error)
}

func NewCateGoryDataService(cateService *repository.CategoryRespository)ICategoryDataService{
	return &CategoryDataService{
		cate: cateService,
	}
}

type CategoryDataService struct {
	cate repository.ICateGoryRepository
}

func (c *CategoryDataService) FindCategoryByName(name string) (*model.Category, error) {
	return c.cate.FindCategoryByName(name)
}

func (c *CategoryDataService) FindCategoryByLevel(levelid uint32) ([]*model.Category, error) {
	return c.cate.FindCategoryByLevel(levelid)
}

func (c *CategoryDataService) FindCategoryByParent(parentId int64) ([]*model.Category, error) {
	return c.cate.FindCategoryByParent(parentId)
}

func (c *CategoryDataService) AddCategory(cate *model.Category) (int64, error) {
	return c.cate.CreateCategory(cate)
}

func (c *CategoryDataService) DeleteCategory(id int64) error {
	return c.cate.DeleteCategoryById(id)
}

func (c *CategoryDataService) UpdateCategory(cate *model.Category) error {
	return c.cate.UpdateCategory(cate)
}

func (c *CategoryDataService) FindCategoryById(id int64) (*model.Category, error) {
	return c.cate.FindCategoryById(id)
}

func (c *CategoryDataService) FindAllCategory() ([]*model.Category, error) {
	return c.cate.FindALl()
}

