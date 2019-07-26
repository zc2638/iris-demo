package model

type User struct {
	AutoID
	UID        string `json:"uid"`                        // 用户id
	Name       string `gorm:"size:50" json:"name"`        // 名称
	Gender     uint   `gorm:"type:tinyint" json:"gender"` // 性别（1男，2女）
	JobNumber  string `json:"job_number"`                 // 工号
	Summary    string `json:"summary"`                    // 业务职能
	Role       string `json:"role"`                       // 系统角色
	Department string `json:"department"`                 // 所属部门
	FaceToken  string `json:"face_token"`                 // 人脸识别face++ token
	FaceImage  string `json:"face_image"`                 // 照片
	Status     uint   `gorm:"type:tinyint" json:"status"` // 状态
	Timestamps
}
