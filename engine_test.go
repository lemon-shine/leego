package leego

import (
	"testing"
)

func TestEngine(t *testing.T) {
	router := NewEngine()

	router.GET("/get", func(ctx *Context) {
		t.Log(ctx.Method, ctx.Path)
		ctx.ResponseBytes(200, []byte("/get"))
	})
	router.POST("/post", func(ctx *Context) {
		t.Log(ctx.Method, ctx.Path)
		ctx.ResponseFormatString(200, "/post/%s", ctx.Query("name"))
	})

	admin := router.Group("/admin")
	admin.PUT("/put", func(ctx *Context) {
		t.Log(ctx.Method, ctx.Path)
		ctx.ResponseJson(200, J{
			"name": ctx.GetPostForm("name"),
			"age":  ctx.GetPostForm("age"),
		})
	})
	admin.OPTIONS("/options", func(ctx *Context) {
		t.Log(ctx.Method, ctx.Path)
		ctx.ResponseString(200, "/options")
	})

	router.ListenAndServe(":8080")
}
