package service

import (
	"errors"
	"sop/data"
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
	db.Where("product_id = ?", id).Find(&all)
	return
}

// 根据id获取产品
func (s *ProductService) GetProductByID(id interface{}) (product model.Product) {

	db := database.NewDB()
	db.Where("id = ?", id).First(&product)
	return
}

// 根据产品名获取产品
func (s *ProductService) GetProductByName(name string) (product model.Product) {

	db := database.NewDB()
	db.Where("name = ?", name).First(&product)
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

// 根据产品型号名称获取产品型号
func (s *ProductService) GetModelByName(productID interface{}, modelName string) (model model.ProductModel) {

	db := database.NewDB()
	db.Where("product_id = ?", productID).Where("name = ?", modelName).First(&model)
	return
}

// 根据产品型号id集合 获取产品型号
func (s *ProductService) GetModelByIDs(ids interface{}) (set []model.ProductModel) {

	db := database.NewDB()
	db.Where("id in (?)", ids).Preload("Product").Find(&set)
	return
}

// 批量创建
func (s *ProductService) Insert(products []data.PartyProduct) error {

	db := database.NewDB()
	tx := db.Begin()
	for _, p := range products {
		if p.Model == nil || len(p.Model) == 0 {
			continue
		}
		product := model.Product{Name: p.Name}
		tx.Create(&product)
		if tx.NewRecord(product) == true {
			tx.Rollback()
			return errors.New("创建失败")
		}
		for _, m := range p.Model {
			productModel := model.ProductModel{
				Name: m.Name,
				ProductID: product.ID,
			}
			tx.Create(&productModel)
			if tx.NewRecord(productModel) == true {
				tx.Rollback()
				return errors.New("创建失败")
			}
		}
	}
	tx.Commit()
	return nil
}
