package model

import (
	"time"
)

type Andon struct {
	AutoID
	Type        string    `json:"type"`         // 类型
	Content     string    `json:"content"`      // 内容
	Station     string    `json:"station"`      // 工位
	StationName string    `json:"station_name"` // 工位名称
	Code        string    `json:"code"`         // 业务号
	Information string    `json:"information"`  // 业务信息
	TriggerTime time.Time `json:"trigger_time"` // 触发时间
	Timestamps
}
