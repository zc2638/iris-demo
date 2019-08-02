package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/errors"
	"strconv"
)

type Base struct {
	Ctx  iris.Context
	Data map[string]interface{}
}

func (c *Base) Err(msg string) {

	fmt.Println("[Error]:", msg)
	_, err := c.Ctx.JSON(iris.Map{
		"code":    400,
		"message": msg,
	})

	if err != nil {
		c.Ctx.StatusCode(iris.StatusInternalServerError)
	}
}

func (c *Base) Succ(msg string, data ...interface{}) {

	m := iris.Map{
		"code":    200,
		"message": msg,
		"data":    nil,
	}

	if data != nil && len(data) > 0 {
		m["data"] = data[0]
	}

	if _, err := c.Ctx.JSON(m); err != nil {
		c.Ctx.StatusCode(iris.StatusInternalServerError)
	}
}

func (c *Base) Text(str string) {
	_, err := c.Ctx.Text(str)
	if err != nil {
		c.Err(err.Error())
	}
}

func (c *Base) Json(v interface{}, options ...context.JSON) {
	_, err := c.Ctx.JSON(v, options...)
	if err != nil {
		c.Err(err.Error())
	}
}

func (c *Base) HTML(htmlContent string) {
	_, err := c.Ctx.HTML(htmlContent)
	if err != nil {
		c.Err(err.Error())
	}
}

func (c *Base) GetMeta() (page, pageSize int, err error) {

	p := c.Ctx.URLParamDefault("page", "1")
	page, err = strconv.Atoi(p)
	if err != nil {
		err = errors.New("分页参数类型错误")
		return
	}

	row := c.Ctx.URLParamDefault("pageSize", "10")
	pageSize, err = strconv.Atoi(row)
	if err != nil {
		err = errors.New("分页参数类型错误")
		return
	}
	return
}
