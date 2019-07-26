package service

import (
	"github.com/jinzhu/gorm"
)

type BaseService struct {}

func (s *BaseService) Paginate(db *gorm.DB, page, pageSize int) *gorm.DB {
	return db.Limit(pageSize).Offset((page - 1) * pageSize)
}