package service

import (
	"sop/lib/database"
	"sop/model"
)

type CraftService struct{ BaseService }

// 获取所有工艺方案
func (s *CraftService) GetAll() (all []model.Craft) {

	db := database.NewDB()
	db.Preload("CraftItem").Find(&all)
	return
}

// 根据id获取工艺方案
func (s *CraftService) GetCraftByID(id interface{}) (craft model.Craft) {

	db := database.NewDB()
	db.Where("id = ?", id).Preload("CraftItem").First(&craft)
	return
}

// 根据工序id集合 获取工艺工序
func (s *CraftService) GetItemByIds(ids interface{}) (set []model.CraftItem) {

	db := database.NewDB()
	db.Where("id in (?)", ids).Find(&set)
	return
}