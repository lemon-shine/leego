/*******************************************************************************
模块：引擎
作者：Lemine
时间：2020/03/07
*******************************************************************************/
package leego

import (
	"net/http"
)

//定义请求处理函数
type HandleFunc func(*Context)

//定义引擎，用于处理所有HTTP请求
type Engine struct {
	*RouteGroup               //继承路由组，用于创建分组
	router      *router       //路由器
	groups      []*RouteGroup //路由组
}

func NewEngine() *Engine {
	engine := &Engine{router: newRouter()} //新建一个路由器
	engine.RouteGroup = &RouteGroup{engine: engine}
	engine.groups = []*RouteGroup{engine.RouteGroup} //不同路由组
	return engine
}

//GET：处理GET请求
func (self *Engine) GET(path string, handler HandleFunc) {
	self.router.addRoute("GET", path, handler)
}

//POST：处理POST请求
func (self *Engine) POST(path string, handler HandleFunc) {
	self.router.addRoute("POST", path, handler)
}

//PUT：处理PUT请求
func (self *Engine) PUT(path string, handler HandleFunc) {
	self.router.addRoute("PUT", path, handler)
}

//DELETE：处理DELETE请求
func (self *Engine) DELETE(path string, handler HandleFunc) {
	self.router.addRoute("DELETE", path, handler)
}

//PATCH：PATCH请求
func (self *Engine) PATCH(path string, handler HandleFunc) {
	self.router.addRoute("PATCH", path, handler)
}

//HEAD：处理HEAD请求
func (self *Engine) HEAD(path string, handler HandleFunc) {
	self.router.addRoute("HEAD", path, handler)
}

//OPTIONS：处理OPTIONS请求
func (self *Engine) OPTIONS(path string, handler HandleFunc) {
	self.router.addRoute("OPTIONS", path, handler)
}

//Run：运行服务引擎
func (self *Engine) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, self)
}

//ServeHTTP：实现http.Hanlder接口，处理所有的路由请求
func (self *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := NewContext(resp, req)
	self.router.handle(ctx)
}
