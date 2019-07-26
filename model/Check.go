package model

// AI识别所需信息
type Check struct {
	AutoID
	Url    string `json:"url"`    // 图片url
	Colors string `json:"colors"` // 颜色信息
	Size   string `json:"size"`   // 图片尺寸
}
