/*******************************************************************************
模块：引擎模块
作者：Lemine
时间：2020/03/01
*******************************************************************************/
package leego

import (
	"fmt"
	"net/http"
)

//定义请求处理函数
type HandleFunc func(http.ResponseWriter, *http.Request)

//定义引擎，用于处理请求
type Engine struct {
	router map[string]HandleFunc
}

func NewEngine() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

//addRoute：添加请求路由到路由器
func (self *Engine) addRoute(method string, path string, handler HandleFunc) {
	key := method + "-" + path
	self.router[key] = handler
}

//GET：处理GET请求
func (self *Engine) GET(path string, handler HandleFunc) {
	self.addRoute("GET", path, handler)
}

//POST：处理POST请求
func (self *Engine) POST(path string, handler HandleFunc) {
	self.addRoute("POST", path, handler)
}

//PUT：处理PUT请求
func (self *Engine) PUT(path string, handler HandleFunc) {
	self.addRoute("PUT", path, handler)
}

//DELETE：处理DELETE请求
func (self *Engine) DELETE(path string, handler HandleFunc) {
	self.addRoute("DELETE", path, handler)
}

//PATCH：PATCH请求
func (self *Engine) PATCH(path string, handler HandleFunc) {
	self.addRoute("PATCH", path, handler)
}

//HEAD：处理HEAD请求
func (self *Engine) HEAD(path string, handler HandleFunc) {
	self.addRoute("HEAD", path, handler)
}

//OPTIONS：处理OPTIONS请求
func (self *Engine) OPTIONS(path string, handler HandleFunc) {
	self.addRoute("OPTIONS", path, handler)
}

//Run：运行服务引擎
func (self *Engine) ListenAndServe(addr string) {
	http.ListenAndServe(addr, self)
}

//ServeHTTP：实现http.Hanlder接口，处理路由请求
func (self *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	//获取请求URL
	key := req.Method + "-" + req.URL.Path
	//判断请求URL对应的处理函数是否存在
	if handler, ok := self.router[key]; ok {
		//处理请求
		handler(resp, req)
	} else {
		fmt.Fprintf(resp, "404 NOT FOUND: %s\n", req.URL)
	}
}
