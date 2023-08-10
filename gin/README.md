从官方网站最简单的例子入手 https://github.com/gin-gonic/gin

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

```

`gin.Default() `获取结构体 `Engine`，代表一个web服务器。一个go进程可获取多个 `Engine`。

`Engine` 组合了 `RouterGroup`，是实现路由功能的核心组件。

关注 `gin.Context`。

测试四种路由方式：

+ 静态路由
+ 参数路由
+ 通配符路由
+ 查询参数





