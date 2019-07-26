package model

import "time"

// 作业计划
type Aps struct {
	AutoID
	JobPlanNumber string       `json:"job_plan_number"`            // 作业计划号
	SerialNo      string       `json:"serial_no"`                  // 产线编号
	ModelID       uint         `json:"model_id"`                   // 产品型号id
	CraftID       uint         `json:"craft_id"`                   // 工艺id
	SopID         uint         `json:"sop_id"`                     // sop标识
	PlanTotal     int          `json:"plan_total"`                 // 计划加工数量
	PlanNum       int          `json:"plan_num"`                   // 计划完工数量
	CompleteNum   int          `json:"complete_num"`               // 实际完工数量
	Status        int          `gorm:"type:tinyint" json:"status"` // 状态
	ApsOrder      []ApsOrder   `gorm:"ForeignKey:ApsID"`
	Sop           Sop          `gorm:"ForeignKey:ID;AssociationForeignKey:SopID"`
	ProductModel  ProductModel `gorm:"ForeignKey:ID;AssociationForeignKey:ModelID"`
	Timestamps
}

// 作业计划-工单
type ApsOrder struct {
	AutoID
	OrderID         string            `json:"order_id"`                       // 工单号(外部系统维护)
	ApsID           uint              `json:"aps_id"`                         // 作业计划id
	Uid             uint              `json:"uid"`                            // 用户id
	ProcessID       uint              `json:"process_id"`                     // 工艺工序id
	SopProcessID    uint              `json:"sop_process_id"`                 // sop工序id
	Station         uint              `json:"station"`                        // 工位号
	StationName     string            `json:"station_name"`                   // 工位名称
	Total           int               `json:"total"`                          // 任务加工数量
	Num             int               `json:"num"`                            // 当前加工数量
	CompleteNum     int               `json:"complete_num"`                   // 实际加工数量
	Status          uint              `gorm:"type:tinyint" json:"status"`     // 状态
	StartAt         time.Time         `gorm:"type:timestamp" json:"start_at"` // 开始时间
	EndAt           time.Time         `gorm:"type:timestamp" json:"end_at"`   // 结束时间
	SopProcess      SopProcess        `gorm:"ForeignKey:ID;AssociationForeignKey:SopProcessID"`
	Aps             Aps               `gorm:"ForeignKey:ID;AssociationForeignKey:ApsID"`
	ApsOrderQuality []ApsOrderQuality `gorm:"ForeignKey:OrderID"`
	Timestamps
}

// 工单质检表
type ApsOrderQuality struct {
	AutoID
	OrderID uint   `json:"order_id"` // 工单id
	PieceNo string `json:"piece_no"` // 工件编号
	Result  string `json:"result"`   // 质检结果
	Remark  string `json:"remark"`   // 备注
}
