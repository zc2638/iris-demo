package service

import (
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