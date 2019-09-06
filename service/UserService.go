package service

import (
	"github.com/kataras/iris/core/errors"
	"sop/lib/database"
	"sop/model"
)

type UserService struct{ BaseService }

// 获取所有用户
func (s *UserService) GetAll() (all []model.User) {

	db := database.NewDB()
	db.Find(&all)
	return
}

// 根据id获取用户
func (s *UserService) GetUserByID(id interface{}) (user model.User) {

	db := database.NewDB()
	db.Where("id = ?", id).First(&user)
	return
}

// 根据uid获取用户
func (s *UserService) GetUserByUid(uid interface{}) (user model.User) {

	db := database.NewDB()
	db.Where("uid = ?", uid).First(&user)
	return
}

// 根据faceToken获取用户
func (s *UserService) GetUserByFaceToken(faceToken string) (user model.User) {

	db := database.NewDB()
	db.Where("face_token = ?", faceToken).First(&user)
	return
}

// 更新用户
func (s *UserService) UpdateOne(user model.User) int64 {

	db := database.NewDB()
	return db.Save(&user).RowsAffected
}

// 批量创建
func (s *UserService) Insert(users []model.User) error {

	db := database.NewDB()
	tx := db.Begin()
	for _, user := range users {
		tx.Create(&user)
		if tx.NewRecord(user) == true {
			tx.Rollback()
			return errors.New("创建失败")
		}
	}
	tx.Commit()
	return nil
}