/*******************************************************************************
模块：路由器
作者：Lemine
时间：2020/03/03
*******************************************************************************/
package leego

import (
	"log"
	"net/http"
)

//定义路由器
type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandleFunc),
	}
}

//addRoute：添加请求路由到路由器中
func (self *router) addRoute(method, path string, handler HandleFunc) {
	log.Printf("Route %4s - %s", method, path)
	key := method + "-" + path
	self.handlers[key] = handler
}

//hanlde：通过上下文处理所有的HTTP请求
func (self *router) handle(ctx *Context) {
	key := ctx.Method + "-" + ctx.Path

	//判断请求URL对应的处理函数是否存在
	if handler, ok := self.handlers[key]; ok {
		handler(ctx)
	} else {
		ctx.ResponseFormatString(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
}
