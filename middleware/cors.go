package middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"net/http"
)

func Cors() context.Handler {
	return func(c iris.Context) {
		method := c.Method()

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, X-Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.StatusCode(http.StatusNoContent)
			return
		}
		// 处理请求
		c.Next()
	}
}