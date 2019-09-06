package service

import (
	"github.com/kataras/iris/core/errors"
	"sop/data"
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

// 根据工艺名称获取
func (s *CraftService) GetCraftByName(name string) (craft model.Craft) {

	db := database.NewDB()
	db.Where("name = ?", name).Preload("CraftItem").First(&craft)
	return
}

// 根据工序id集合 获取工艺工序
func (s *CraftService) GetItemByIds(ids interface{}) (set []model.CraftItem) {

	db := database.NewDB()
	db.Where("id in (?)", ids).Find(&set)
	return
}

// 批量创建
func (s *CraftService) Insert(crafts []data.PartyCraft) error {

	db := database.NewDB()
	tx := db.Begin()
	for _, c := range crafts {
		if c.Item == nil || len(c.Item) == 0 {
			continue
		}
		craft := model.Craft{Name: c.Name}
		tx.Create(&craft)
		if tx.NewRecord(craft) == true {
			tx.Rollback()
			return errors.New("创建失败")
		}
		for _, m := range c.Item {
			item := model.CraftItem{
				CraftID:      craft.ID,
				Name:         m.Name,
				CheckImg:     m.CheckImg,
				MinefieldImg: m.MinefieldImg,
				Sort:         m.Sort,
			}
			tx.Create(&item)
			if tx.NewRecord(item) == true {
				tx.Rollback()
				return errors.New("创建失败")
			}
		}
	}
	tx.Commit()
	return nil
}
