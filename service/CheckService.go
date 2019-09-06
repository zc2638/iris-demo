package service

import (
	"github.com/kataras/iris/core/errors"
	"sop/lib/database"
	"sop/model"
)

type CheckService struct{ BaseService }

// 根据url获取单条数据
func (s *CheckService) GetCheckByUrl(url interface{}) (check model.Check) {

	db := database.NewDB()
	db.Where("url = ?", url).First(&check)
	return
}

// 添加
func (s *CheckService) Create(check model.Check) error {

	db := database.NewDB()
	db.Create(&check)
	if db.NewRecord(check) == true {
		return errors.New("check创建失败")
	}
	return nil
}

// 更新
func (s *CheckService) UpdateOne(check model.Check) error {

	db := database.NewDB()
	if db.Save(&check).RowsAffected != 1 {
		return errors.New("操作失败")
	}
	return nil
}

// 批量创建
func (s *CheckService) Insert(checks []model.Check) error {

	db := database.NewDB()
	tx := db.Begin()
	for _, check := range checks {
		tx.Create(&check)
		if tx.NewRecord(check) == true {
			tx.Rollback()
			return errors.New("创建失败")
		}
	}
	tx.Commit()
	return nil
}