/*******************************************************************************
实现：引擎模块
作者：Lemine
时间：2020/03/03
*******************************************************************************/
package leego

import (
	"net/http"
)

//定义请求处理函数
type HandleFunc func(*Context)

//定义引擎，用于处理所有HTTP请求
type Engine struct {
	router *router
}

func NewEngine() *Engine {
	return &Engine{router: newRouter()}
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

//ListenAndServe：运行服务引擎
func (self *Engine) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, self)
}

//ServeHTTP：实现http.Hanlder接口，处理所有的路由请求
func (self *Engine) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	ctx := NewContext(wr, req)
	self.router.handle(ctx)
}
