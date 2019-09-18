package controller

import (
	"encoding/json"
	"sop/data"
	"sop/model"
	"sop/service"
	"strings"
)

/**
 * Created by zc on 2019-09-04.
 */

type PartyController struct{ Base }

// andon数据同步
func (c *PartyController) PostSyncAndon() {

	source := c.Ctx.PostValue("source")
	trimSource := strings.Replace(source, " ", "", -1)
	if trimSource == "[]" || trimSource == "[{}]" || trimSource == "[{},]" {
		c.Err("请求数据不能为空")
		return
	}

	var andonData []data.PartyAndon
	if err := json.Unmarshal([]byte(source), &andonData); err != nil {
		c.Err("请求格式错误")
		return
	}

	var andons = make([]model.Andon, 0)
	for _, andon := range andonData {
		andons = append(andons, model.Andon{
			Type:        andon.Type,
			Content:     andon.Content,
			Station:     andon.Station,
			StationName: andon.StationName,
			Code:        andon.Code,
			Information: andon.Information,
			TriggerTime: andon.TriggerTime,
		})
	}

	if err := new(service.AndonService).Insert(andons); err != nil {
		c.Err("同步失败")
		return
	}
	c.Succ("同步成功")
}

// 产品/产品型号数据同步
func (c *PartyController) PostSyncProduct() {

	source := c.Ctx.PostValue("source")
	trimSource := strings.Replace(source, " ", "", -1)
	if trimSource == "[]" || trimSource == "[{}]" || trimSource == "[{},]" {
		c.Err("请求数据不能为空")
		return
	}

	var productData []data.PartyProduct
	if err := json.Unmarshal([]byte(source), &productData); err != nil {
		c.Err("请求格式错误")
		return
	}

	if err := new(service.ProductService).Insert(productData); err != nil {
		c.Err("同步失败")
		return
	}
	c.Succ("同步成功")
}

// 工艺方案（工序）同步
func (c *PartyController) PostSyncCraft() {

	source := c.Ctx.PostValue("source")
	trimSource := strings.Replace(source, " ", "", -1)
	if trimSource == "[]" || trimSource == "[{}]" || trimSource == "[{},]" {
		c.Err("请求数据不能为空")
		return
	}

	var craftData []data.PartyCraft
	if err := json.Unmarshal([]byte(source), &craftData); err != nil {
		c.Err("请求格式错误")
		return
	}

	if err := new(service.CraftService).Insert(craftData); err != nil {
		c.Err("同步失败")
		return
	}
	c.Succ("同步成功")
}

// 作业计划同步
func (c *PartyController) PostSyncAps() {

	source := c.Ctx.PostValue("source")
	trimSource := strings.Replace(source, " ", "", -1)
	if trimSource == "[]" || trimSource == "[{}]" || trimSource == "[{},]" {
		c.Err("请求数据不能为空")
		return
	}

	var apsData []data.PartyAps
	if err := json.Unmarshal([]byte(source), &apsData); err != nil {
		c.Err("请求格式错误")
		return
	}
	if err := new(service.ApsService).Insert(apsData); err != nil {
		c.Err("同步失败")
		return
	}
	c.Succ("同步成功")
}

// 工单质检同步
func (c *PartyController) PostSyncQuality() {

	source := c.Ctx.PostValue("source")
	trimSource := strings.Replace(source, " ", "", -1)
	if trimSource == "[]" || trimSource == "[{}]" || trimSource == "[{},]" {
		c.Err("请求数据不能为空")
		return
	}

	var qualityData []data.PartyQuality
	if err := json.Unmarshal([]byte(source), &qualityData); err != nil {
		c.Err("请求格式错误")
		return
	}
	if err := new(service.ApsService).InsertQuality(qualityData); err != nil {
		c.Err("同步失败")
		return
	}
	c.Succ("同步成功")
}

// 用户数据同步
func (c *PartyController) PostSyncUser() {

	source := c.Ctx.PostValue("source")
	trimSource := strings.Replace(source, " ", "", -1)
	if trimSource == "[]" || trimSource == "[{}]" || trimSource == "[{},]" {
		c.Err("请求数据不能为空")
		return
	}

	var userData []struct {
		UID        string `json:"uid"`                        // 用户id
		Name       string `gorm:"size:50" json:"name"`        // 名称
		Gender     uint   `gorm:"type:tinyint" json:"gender"` // 性别（1男，2女）
		JobNumber  string `json:"job_number"`                 // 工号
		Summary    string `json:"summary"`                    // 业务职能
		Role       string `json:"role"`                       // 系统角色
		Department string `json:"department"`                 // 所属部门
	}

	if err := json.Unmarshal([]byte(source), &userData); err != nil {
		c.Err("请求格式错误")
		return
	}

	var users = make([]model.User, 0)
	for _, user := range userData {
		users = append(users, model.User{
			UID:        user.UID,
			Name:       user.Name,
			Gender:     user.Gender,
			JobNumber:  user.JobNumber,
			Summary:    user.Summary,
			Role:       user.Role,
			Department: user.Department,
		})
	}

	if err := new(service.UserService).Insert(users); err != nil {
		c.Err("同步失败")
		return
	}
	c.Succ("同步成功")
}

// POKA-YOKE数据同步
func (c *PartyController) PostSyncCheck() {

	source := c.Ctx.PostValue("source")
	trimSource := strings.Replace(source, " ", "", -1)
	if trimSource == "[]" || trimSource == "[{}]" || trimSource == "[{},]" {
		c.Err("请求数据不能为空")
		return
	}

	var checkData []struct {
		Url    string `json:"url"`    // 图片url
		Colors string `json:"colors"` // 颜色信息
		Size   string `json:"size"`   // 图片尺寸
	}
	if err := json.Unmarshal([]byte(source), &checkData); err != nil {
		c.Err("请求格式错误")
		return
	}

	var checks = make([]model.Check, 0)
	for _, check := range checkData {
		checks = append(checks, model.Check{
			Url:    check.Url,
			Colors: check.Colors,
			Size:   check.Size,
		})
	}

	if err := new(service.CheckService).Insert(checks); err != nil {
		c.Err("同步失败")
		return
	}
	c.Succ("同步成功")
}
