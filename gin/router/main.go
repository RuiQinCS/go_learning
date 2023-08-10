package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// 静态路由
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "静态路由")
	})

	// 参数路由
	server.GET("/users/:name", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "参数路由。参数：%v", ctx.Param("name"))
	})

	//通配符路由
	server.GET("/views/*.html", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "通配符路由。匹配的值为：%s", ctx.Param(".html"))
	})

	// 查询参数
	server.GET("/order", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "查询参数id为:%v", ctx.Query("id"))
	})

	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
