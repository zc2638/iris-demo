package controller

import "sop/service"

type UserController struct{ Base }

// 用户列表（不分页）
func (c *UserController) GetList() {

	all := new(service.UserService).GetAll()

	var data = make([]map[string]interface{}, 0)
	for _, v := range all {
		data = append(data, map[string]interface{}{
			"id":         v.ID,         // 用户id
			"name":       v.Name,       // 用户名称
			"gender":     v.Gender,     // 性别
			"image":      v.FaceImage,  // 人脸识别照
			"jobNumber":  v.JobNumber,  // 工号
			"summary":    v.Summary,    // 业务职能
			"role":       v.Role,       // 系统角色
			"department": v.Department, // 所属部门
		})
	}

	c.Succ("", data)
}
