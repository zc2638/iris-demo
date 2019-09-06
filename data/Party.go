package data

import "time"

/**
 * Created by zc on 2019-09-04.
 */

type PartyAndon struct {
	Type        string    `json:"type"`         // 类型
	Content     string    `json:"content"`      // 内容
	Station     string    `json:"station"`      // 工位
	StationName string    `json:"station_name"` // 工位名称
	Code        string    `json:"code"`         // 业务号
	Information string    `json:"information"`  // 业务信息
	TriggerTime time.Time `json:"trigger_time"` // 触发时间
}

type PartyProduct struct {
	Name  string `json:"name"`
	Model []struct {
		Name string `json:"name"`
	} `json:"model"`
}

type PartyCraft struct {
	Name string `json:"name"`
	Item []struct {
		Name         string `json:"name"`
		CheckImg     string `json:"check_img"`
		MinefieldImg string `json:"minefield_img"`
		Sort         uint   `json:"sort"`
	} `json:"item"`
}

type PartyAps struct {
	JobPlanNumber string       `json:"job_plan_number"` // 作业计划号
	SerialNo      string       `json:"serial_no"`       // 产线编号
	Product       string       `json:"product"`         // 产品名称
	Model         string       `json:"model"`           // 产品型号
	CraftName     string       `json:"craft_name"`      // 工艺名称
	PlanTotal     int          `json:"plan_total"`      // 计划加工数量
	PlanNum       int          `json:"plan_num"`        // 计划完工数量
	CompleteNum   int          `json:"complete_num"`    // 实际完工数量
	Order         []PartyOrder `json:"order"`
}

type PartyOrder struct {
	OrderID       string    `json:"order_id"`        // 工单号(外部系统维护)
	Uid           uint      `json:"uid"`             // 用户id（外部系统）
	CraftItemName string    `json:"craft_item_name"` // 工艺工序名称
	Station       uint      `json:"station"`         // 工位号
	StationName   string    `json:"station_name"`    // 工位名称
	Total         int       `json:"total"`           // 任务加工数量
	Num           int       `json:"num"`             // 当前加工数量
	CompleteNum   int       `json:"complete_num"`    // 实际加工数量
	StartAt       time.Time `json:"start_at"`        // 开始时间
	EndAt         time.Time `json:"end_at"`          // 结束时间
}

type PartyQuality struct {
	OrderID string `json:"order_id"` // 工单id
	PieceNo string `json:"piece_no"` // 工件编号
	Result  string `json:"result"`   // 质检结果
	Remark  string `json:"remark"`   // 备注
}
