# Leego Web 框架

​		Leego 是一款用 Golang 编写的仿 Gin 的 Web 框架，仅供学习使用。

## 目录

- Leego Web 框架
  - 目录
  - 安装
  - 快速使用
  - API 示例
    - 使用 GET、POST、PUT、DELETE、PATCH、HEAD 和 OPTIONS 请求方法
    - 查询路由中的请求参数
    - 获取表单中的请求参数

## 安装

1. 先安装 **Golang  1.11+ **，然后使用 go get 命令安装：

```shell
$ go get -u github.com/lemon-shine/leego
```

2. 在代码中导入：

```go
import "github.com/lemon-shine/leego"
```

## 快速使用

```go
package main

import (
	"net/http"
	"github.com/lemon-shine/leego"
)

func main() {
	router := leego.Engine()
	router.GET("/ping", func(resp http.ResponseWriter, req *http.Request) {
		resp.SetHeader(200)
		resp.Write(string("/pong"))
	})
	router.ListenAndServe(":8080")
}
```

## API 示例

### 使用 GET、POST、PUT、DELETE、HEAD 和 OPTIONS 请求方法

```go
func main() {
	router := leego.Engine()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)
    
	router.ListenAndServe(":8080")
}
```

### 查询路由中的请求参数

```go
func main() {
	router := leego.Engine()
    
    //请求URL：/index?name=xxx&age=yyy
	router.GET("/index", func(ctx *leego.Context) {
		name := ctx.Query("name")
        age := ctx.Query("age")
        ctx.ResponseString(http.StatusOK, "hello " + name)
	})
	router.ListenAndServe(":8080")
}
```

### 获取表单中的请求参数

```go
func main() {
	router := leego.Engine()

	router.POST("/form", func(ctx *leego.Context) {
		info := ctx.GetPostForm("info")

		ctx.ResponseJson(200, leego.J{
            "name":name,
            "age":age,
		})
	})
	router.ListenAndServe(":8080")
}
```