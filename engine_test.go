package leego

import (
	"testing"
)

func TestEngine(t *testing.T) {
	router := NewEngine()
	router.GET("/", func(ctx *Context) {
		ctx.ResponseString(200, "wellcome")
	})

	group := router.NewGroup("/aaa")
	group.AddMiddlewares(Logger(), Authenticate(&Auth{Username: "123", Password: "123"}))
	group.GET("/bbb", func(ctx *Context) {
		ctx.ResponseString(200, "hello")
	})

	group.GET("/ccc", func(ctx *Context) {
		ctx.ResponseString(200, ctx.Query("username"))
	})

	group.GET("/file/:file_id/download", func(ctx *Context) {
		ctx.ResponseString(200, ctx.GetParam("file_id"))
	})
	router.ListenAndServe(":8080")
}
