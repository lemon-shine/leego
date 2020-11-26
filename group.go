/*******************************************************************************
模块：路由组
作者：Lemine
时间：2020/03/07
*******************************************************************************/
package leego

//定义路由组
type RouteGroup struct {
	prefix      string       //路由组名
	middlewares []HandleFunc //路由组所支持的中间件
	engine      *Engine      //绑定路由组到引擎
}

func (self *RouteGroup) Group(prefix string) *RouteGroup {
	engine := self.engine
	group := &RouteGroup{
		prefix: self.prefix + prefix, //新路由组名
		engine: engine,
	}
	engine.groups = append(engine.groups, group) //添加新路由组到引擎

	return group
}

//addRoute：添加新路由到指当前路由组
func (self *RouteGroup) addRoute(method string, part string, handler HandleFunc) {
	path := self.prefix + part
	self.engine.router.addRoute(method, path, handler)
}

//GET：处理当前路由组的GET请求
func (self *RouteGroup) GET(part string, handler HandleFunc) {
	self.addRoute("GET", part, handler)
}

//POST：处理当前路由组的POST请求
func (self *RouteGroup) POST(part string, handler HandleFunc) {
	self.addRoute("POST", part, handler)
}

//PUT：处理当前路由组的PUT请求
func (self *RouteGroup) PUT(part string, handler HandleFunc) {
	self.addRoute("PUT", part, handler)
}

//DELETE：处理当前路由组的DELETE请求
func (self *RouteGroup) DELETE(part string, handler HandleFunc) {
	self.addRoute("DELETE", part, handler)
}

//PATCH：处理当前路由组的PATCH请求
func (self *RouteGroup) PATCH(path string, handler HandleFunc) {
	self.addRoute("PATCH", path, handler)
}

//HEAD：处理当前路由组的HEAD请求
func (self *RouteGroup) HEAD(path string, handler HandleFunc) {
	self.addRoute("HEAD", path, handler)
}

//OPTIONS：处理当前路由组的OPTIONS请求
func (self *RouteGroup) OPTIONS(path string, handler HandleFunc) {
	self.addRoute("OPTIONS", path, handler)
}
