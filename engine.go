/*******************************************************************************
Method: HTTP/HTTPS引擎
Author: Lemine
Langua: Golang 1.14
Modify：2020/03/07
*******************************************************************************/
package leego

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

const (
	//引擎默认支持的最大请求头
	defaultMaxHeaderBytes = 10 << 20 //10MB
	//引擎默认支持的最大请求体
	defaultMaxBodyBytes = 30 << 20 //30MB
)

//定义请求处理函数
type HandleFunc func(*Context)

//定义引擎
type Engine struct {
	//自带路由组
	*RouteGroup
	//路由器
	router *router
	//自定义路由组
	groups []*RouteGroup

	//HTML模板
	htmlTemplates *template.Template
	htmlFuncMap   template.FuncMap

	maxBodyBytes int64
}

//不带任何中间件的新引擎
func NewEngine() *Engine {
	//新建一个路由器
	engine := &Engine{router: newRouter()}
	//引擎所属的自带路由组
	engine.RouteGroup = &RouteGroup{engine: engine}
	//将默认路由组添加到路由组
	engine.groups = []*RouteGroup{engine.RouteGroup}
	engine.maxBodyBytes = defaultMaxBodyBytes
	return engine
}

//自带错误恢复中间的件默认引擎
func NewDefaultEngine() *Engine {
	engine := NewEngine()
	engine.AddMiddlewares(Recover())
	return engine
}

//设置HTML的模板函数
func (self *Engine) SetHtmlFuncMap(funcMap template.FuncMap) {
	self.htmlFuncMap = funcMap
}

//
func (self *Engine) LoadHtmlGlob(pattern string) {
	self.htmlTemplates = template.Must(template.New("").Funcs(self.htmlFuncMap).ParseGlob(pattern))
}

//处理GET请求
func (self *Engine) GET(path string, handler HandleFunc) {
	self.router.addRoute("GET", path, handler)
}

//处理POST请求
func (self *Engine) POST(path string, handler HandleFunc) {
	self.router.addRoute("POST", path, handler)
}

//处理PUT请求
func (self *Engine) PUT(path string, handler HandleFunc) {
	self.router.addRoute("PUT", path, handler)
}

//处理DELETE请求
func (self *Engine) DELETE(path string, handler HandleFunc) {
	self.router.addRoute("DELETE", path, handler)
}

//处理PATCH请求
func (self *Engine) PATCH(path string, handler HandleFunc) {
	self.router.addRoute("PATCH", path, handler)
}

//处理HEAD请求
func (self *Engine) HEAD(path string, handler HandleFunc) {
	self.router.addRoute("HEAD", path, handler)
}

//处理OPTIONS请求
func (self *Engine) OPTIONS(path string, handler HandleFunc) {
	self.router.addRoute("OPTIONS", path, handler)
}

//运行HTTP引擎
func (self *Engine) ListenAndServe(addr string) {
	log.Printf("Listening and Serving HTTP server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, self))
}

//运行HTTPS引擎
func (self *Engine) ListenAndServeTLS(addr, certFile, keyFile string) {
	log.Printf("Listening and Serving HTTPS server on %s\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, certFile, keyFile, self))
}

//实现http.Hanlder接口，处理所有的路由请求
func (self *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	var middlewares []HandleFunc
	//将引擎自带路由组的中间件添加到middlewares
	if self.RouteGroup.middlewares != nil {
		middlewares = append(middlewares, self.RouteGroup.middlewares...)
	}

	//遍历寻找匹配的自定义路由组
	for _, group := range self.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			//将匹配成功路由组的中间件放入middlewares
			middlewares = append(middlewares, group.middlewares...)
			// break
		}
	}

	//新建一个上下文
	ctx := newContext(self, req, resp)
	//将middlewares绑定到当前上下文
	ctx.middlewares = middlewares
	self.router.handle(ctx)
}
