package leego

import (
	"testing"
)

func TestEngine(t *testing.T) {
	router := NewEngine()
	router.GET("/get", func(ctx *Context) {
		t.Log(ctx.Method, ctx.Path)
		t.Log(ctx.Query("name"))
		ctx.ResponseBytes(200, []byte(ctx.Path))
	})
	router.POST("/post", func(ctx *Context) {
		t.Log(ctx.Method, ctx.Path)
		ctx.ResponseString(200, "/post/"+ctx.Query("name"))

	})
	router.PUT("/put", func(ctx *Context) {
		ctx.ResponseString(200, "PUT /put")
	})

	router.ListenAndServe(":8080")
}
