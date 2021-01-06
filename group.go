/*******************************************************************************
Method: 路由分组
Author: Lemine
Langua: Golang 1.14
Modify：2020/03/07
*******************************************************************************/
package leego

import (
	"net/http"
	"path"
)

//定义路由组
type RouteGroup struct {
	//路由组名
	//NOTE:路由组名中不能有动态参数
	//NOTE:支持多个路由组嵌套，如：NewGroup("/index").NewGroup("/user").NewGroup("/info")
	prefix string
	//路由组所支持的中间件
	middlewares []HandleFunc
	//路由组所属引擎
	engine *Engine
}

func (self *RouteGroup) NewGroup(prefix string) *RouteGroup {
	engine := self.engine
	group := &RouteGroup{
		prefix: self.prefix + prefix, //新路由组名(前缀)
		engine: engine,
	}
	engine.groups = append(engine.groups, group)

	return group
}

//添加中间件到当前路由组
func (self *RouteGroup) AddMiddlewares(middlewares ...HandleFunc) {
	self.middlewares = append(self.middlewares, middlewares...)
}

//处理GET请求
func (self *RouteGroup) GET(path string, handler HandleFunc) {
	self.addRoute("GET", path, handler)
}

//处理POST请求
func (self *RouteGroup) POST(path string, handler HandleFunc) {
	self.addRoute("POST", path, handler)
}

//处理PUT请求
func (self *RouteGroup) PUT(path string, handler HandleFunc) {
	self.addRoute("PUT", path, handler)
}

//处理DELETE请求
func (self *RouteGroup) DELETE(path string, handler HandleFunc) {
	self.addRoute("DELETE", path, handler)
}

//处理PATCH请求
func (self *RouteGroup) PATCH(path string, handler HandleFunc) {
	self.addRoute("PATCH", path, handler)
}

//处理HEAD请求
func (self *RouteGroup) HEAD(path string, handler HandleFunc) {
	self.addRoute("HEAD", path, handler)
}

//处理OPTIONS请求
func (self *RouteGroup) OPTIONS(path string, handler HandleFunc) {
	self.addRoute("OPTIONS", path, handler)
}

//设置静态文件处理器
func (self *RouteGroup) SetStatic(urlPart string, filesPath string) {
	handler := self.createStaticHandler(urlPart, http.Dir(filesPath))
	url := path.Join(urlPart, "/*filepath")
	//使用GET方法请求静态文件
	self.GET(url, handler)
}

//创建当前路由组的静态文件处理器
func (self *RouteGroup) createStaticHandler(urlPart string, fs http.FileSystem) HandleFunc {
	//拼接绝对路径
	absolutePath := path.Join(self.prefix, urlPart)
	//去除静态文件处理器的路由前缀
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(ctx *Context) {
		//静态文件的文件夹命名为：filepath
		file := ctx.GetParam("filepath")
		//打开文件
		if _, err := fs.Open(file); err != nil {
			ctx.SetStatus(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(ctx.Response, ctx.Request)
	}
}

//添加请求路由到当前路由组
func (self *RouteGroup) addRoute(method string, part string, handler HandleFunc) {
	path := self.prefix + part
	self.engine.router.addRoute(method, path, handler)
}
