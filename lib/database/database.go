package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sop/model"
	"time"
)

var DB *gorm.DB

func init() {

	var err error

	//DB, err = gorm.Open("mysql", "root:@tcp(localhost:3306)/sop?charset=utf8mb4&parseTime=True&loc=Local")
	//if err != nil {
	//	panic(err)
	//}

	DB, err = gorm.Open("sqlite3", "sop.db")
	if err != nil {
		panic("连接数据库失败")
	}

	DB.DB().SetMaxIdleConns(10)               // 连接池的空闲数大小
	DB.DB().SetMaxOpenConns(100)              // 最大打开连接数
	DB.DB().SetConnMaxLifetime(time.Hour * 6) // 连接最长存活时间

	// 结构生成
	DBMigrate()
	// 数据填充
	seed()
}

func NewDB() *gorm.DB {
	return DB
}

func DBMigrate() {

	// 禁用表名复数
	//DB.SingularTable(true)

	// 自动生成表结构
	DB.AutoMigrate(
		new(model.Admin), new(model.Aps), new(model.ApsOrder), new(model.ApsOrderQuality), new(model.Andon),
		new(model.Craft), new(model.CraftItem),
		new(model.Product), new(model.ProductModel),
		new(model.Sop), new(model.SopModel), new(model.SopProcess),
		new(model.User),
		new(model.Check),
	)
}
