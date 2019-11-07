package model

type Admin struct {
	AutoID
	Name     string `gorm:"size:50;not null;" json:"name"` // 名称
	Password string `gorm:"not null;" json:"password"`     // 密码
	Salt     string `gorm:"size:10;not null;" json:"salt"` // 盐值
	Status   uint   `gorm:"type:tinyint" json:"status"`    // 状态
	Timestamps
}
