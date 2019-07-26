package model

// 工艺
type Craft struct {
	AutoID
	Name      string      `json:"name"`                       // 工艺名称
	Status    uint        `gorm:"type:tinyint" json:"status"` // 状态
	CraftItem []CraftItem `gorm:"ForeignKey:CraftID"`
	Timestamps
}

// 工艺工序
type CraftItem struct {
	AutoID
	CraftID  uint   `json:"craft_id"`                   // 工艺方案id
	Name     string `json:"name"`                       // 工艺工序名称
	CheckImg string `json:"check_img"`                  // 防差错图片
	Sort     uint   `json:"sort"`                       // 工序序号(步骤号)
	Status   uint   `gorm:"type:tinyint" json:"status"` // 状态
	Timestamps
}
