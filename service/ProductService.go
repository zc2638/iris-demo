package service

import (
	"sop/lib/database"
	"sop/model"
)

type ProductService struct{ BaseService }

// 获取所有产品
func (s *ProductService) GetAll() (all []model.Product) {

	db := database.NewDB()
	db.Find(&all)
	return
}

// 根据产品id获取所有产品型号
func (s *ProductService) GetModelAllByProductID(id interface{}) (all []model.ProductModel) {

	db := database.NewDB()
	db.Where("product_id = ?", id).Where("status != ?", model.APS_STATUS_START).Find(&all)
	return
}

// 根据id获取产品
func (s *ProductService) GetProductByID(id interface{}) (product model.Product) {

	db := database.NewDB()
	db.Where("id = ?", id).First(&product)
	return
}

// 根据id获取产品集合
func (s *ProductService) GetProductListByIds(ids interface{}) (list []model.Product) {

	db := database.NewDB()
	db.Where("id in (?)", ids).Find(&list)
	return
}

// 根据id获取产品型号
func (s *ProductService) GetModelByID(id interface{}) (model model.ProductModel) {

	db := database.NewDB()
	db.Where("id = ?", id).Preload("Product").First(&model)
	return
}

// 根据产品型号id集合 获取产品型号
func (s *ProductService) GetModelByIDs(ids interface{}) (set []model.ProductModel) {

	db := database.NewDB()
	db.Where("id in (?)", ids).Preload("Product").Find(&set)
	return
}