package handler

import (
	"github.com/idcpj/micro/common"
	"github.com/idcpj/micro/category/domain/model"
	"github.com/idcpj/micro/category/domain/service"
	"github.com/idcpj/micro/category/proto/category"
	"context"
)

type Category struct {
	CateDataServer service.ICategoryDataService
}

func NewCategory(cateDataServer service.ICategoryDataService) *Category {
	return &Category{CateDataServer: cateDataServer}
}

func (c *Category) CreateCategory(ctx context.Context, request *category.CategoryRequest, response *category.CreateCategoryResponse) error {
	// 手动转换
	//cate :=&model.Category{
	//	CategoryName:        request.GetCategoryName(),
	//	CategoryLevel:       request.GetCategoryLevel(),
	//	CategoryParent:      request.GetCategoryParent(),
	//	CategoryImage:       request.GetCategoryImage(),
	//	CategoryDescription: request.GetCategoryDescription(),
	//}

	// 自动转换
	cate := &model.Category{}
	err := common.SwapTo(request, cate)
	if err != nil {
		return err
	}
	_, err = c.CateDataServer.AddCategory(cate)
	if err != nil {
		return err
	}
	response.CategoryId = cate.ID
	response.Message = "分类添加成功"
	return nil
}

func (c *Category) UpdateCategory(ctx context.Context, request *category.CategoryRequest, response *category.UpdateCategoryResponse) error {
	cate := &model.Category{}
	err := common.SwapTo(request, cate)
	if err != nil {
		return err
	}

	err = c.CateDataServer.UpdateCategory(cate)
	if err != nil {
		return err
	}
	response.Message = "分类更新服务"
	return nil
}

func (c *Category) DeleteCategory(ctx context.Context, request *category.DeleteCategoryRequest, response *category.DeleteCategoryResponse) error {
	cate := &model.Category{}
	err := common.SwapTo(request, cate)
	if err != nil {
		return err
	}

	err = c.CateDataServer.DeleteCategory(request.CategoryId)
	if err != nil {
		return err
	}
	response.Message = "删除成功"
	return nil
}

func (c *Category) FindCategoryByName(ctx context.Context, request *category.FindByNameRequest, response *category.CategoryResponse) error {
	cate, err := c.CateDataServer.FindCategoryByName(request.GetCategoryName())
	if err != nil {
		return err
	}
	return common.SwapTo(cate, response)
}

func (c *Category) FindCategoryByID(ctx context.Context, request *category.FindByIdRequest, response *category.CategoryResponse) error {
	cate, err := c.CateDataServer.FindCategoryById(request.Id)
	if err != nil {
		return err
	}
	return common.SwapTo(cate, response)
}

func (c *Category) FindCategoryByLevel(ctx context.Context, request *category.FindByLevelRequest, response *category.FindAllResponse) error {

	cate, err := c.CateDataServer.FindCategoryByLevel(request.GetCategoryLevel())
	if err != nil {
		return err
	}

	return common.SwapTo(cate, &response.Category)
}

func (c *Category) FindCategoryByParent(ctx context.Context, request *category.FindByParentRequest, response *category.FindAllResponse) error {
	cate, err := c.CateDataServer.FindCategoryByParent(request.GetParentId())
	if err != nil {
		return err
	}
	return common.SwapTo(cate, &response.Category)
}

func (c *Category) FindAllCategory(ctx context.Context, request *category.FindAllRequest, response *category.FindAllResponse) error {
	cates, err := c.CateDataServer.FindAllCategory()
	if err != nil {
		return err
	}
	return common.SwapTo(cates, &response.Category)
}
