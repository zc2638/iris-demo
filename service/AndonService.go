package service

import (
	"errors"
	"sop/lib/database"
	"sop/model"
)

type AndonService struct{ BaseService }

// 获取所有andon
func (s *AndonService) GetAll() (all []model.Andon) {

	db := database.NewDB()
	db.Find(&all)
	return
}

// 创建数据
func (s *AndonService) Create(andon model.Andon) error {

	db := database.NewDB()
	db.Create(&andon)
	if db.NewRecord(andon) == true {
		return errors.New("创建失败")
	}
	return nil
}

// 批量创建数据
func (s *AndonService) Insert(andons []model.Andon) error {

	db := database.NewDB()
	tx := db.Begin()
	for _, andon := range andons {
		tx.Create(&andon)
		if tx.NewRecord(andon) == true {
			tx.Rollback()
			return errors.New("创建失败")
		}
	}
	tx.Commit()
	return nil
}