package model

import "time"

type Sop struct {
	AutoID
	Title      string       `json:"title"`                           // sop标题
	CraftID    uint         `json:"craft_id"`                        // 工艺方案id
	ProductID  uint         `json:"product_id"`                      // 产品id
	IsTemplate uint         `gorm:"type:tinyint" json:"is_template"` // 类型：是否模板
	Version    uint         `json:"version"`                         // 版本号
	Status     uint         `gorm:"type:tinyint" json:"status"`      // 状态
	Comment    string       `json:"comment"`                         // 备注
	PassAt     time.Time    `gorm:"type:timestamp" json:"pass_at"`   // 通过时间
	RefuseAt   time.Time    `gorm:"type:timestamp" json:"refuse_at"` // 驳回时间
	Craft      Craft        `gorm:"ForeignKey:ID;AssociationForeignKey:CraftID"`
	Product    Product      `gorm:"ForeignKey:ID;AssociationForeignKey:ProductID"`
	SopModel   []SopModel   `gorm:"ForeignKey:SopID"`
	SopProcess []SopProcess `gorm:"ForeignKey:SopID"`
	Timestamps
}

type SopModel struct {
	AutoID
	SopID   uint   `json:"sop_id"`   // sop标识
	ModelID uint   `json:"model_id"` // 产品型号id
	Name    string `json:"name"`     // 产品型号名称
	Timestamps
}

type SopProcess struct {
	AutoID
	SopID        uint   `json:"sop_id"`                       //
	ProcessID    uint   `json:"process_id"`                   // 工艺工序id
	Title        string `json:"title"`                        // 工序名称
	Content      string `gorm:"type:text" json:"content"`     // 内容
	Imgs         string `gorm:"type:text" json:"imgs"`        // 内容 转 图片
	Annex        string `json:"annex"`                        // 附件地址（多个）
	CheckImg     string `json:"check_img"`                    // 防差错图片(json数组)
	MinefieldImg string `json:"minefield_img"`                // 雷区预警图片(json数组)
	IsCheck      uint   `gorm:"type:tinyint" json:"is_check"` // 是否开启防差错
	Sort         uint   `json:"sort"`                         // 工序序号（步骤号）
	Timestamps
}
