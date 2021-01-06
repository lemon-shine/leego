/*******************************************************************************
Method: 路由器
Author: Lemine
Langua: Golang 1.14
Modify：2020/03/07
*******************************************************************************/
package leego

import (
	"log"
)

//定义路由器
type router struct {
	//处理器
	handlers map[string]HandleFunc
	//路由树
	tries map[string]*trie
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandleFunc),
		tries:    make(map[string]*trie),
	}
}

//添加路由到路由器中
func (self *router) addRoute(method, path string, handler HandleFunc) {
	log.Printf("Register route %4s - %s", method, path)
	//路由器中不存在该方法，则新建一棵该方法的路由树
	if _, ok := self.tries[method]; !ok {
		self.tries[method] = newTrie()
	}
	self.tries[method].insert(path) //插入请求路由
	key := method + "-" + path      //拼接请求路由
	self.handlers[key] = handler    //绑定请求路由和处理函数
}

//从路由树中获取路由参数
func (self *router) getRoute(method, path string) (string, map[string]string) {
	methodTree, ok := self.tries[method]
	//检查该方法的路由树是否注册
	//self.tries[method].search(path)遇到未注册的方法会panic
	if !ok {
		return "", nil
	}
	node, params := methodTree.search(path) //检索请求路由节点
	if node == nil {
		return "", nil
	}
	path = method + "-" + node.path //拼接请求路由
	return path, params
}

//处理请求上下文
func (self *router) handle(ctx *Context) {
	var path string
	path, ctx.Params = self.getRoute(ctx.Method, ctx.Path) //获取路由和动态参数
	if path == "" {
		ctx.ResponseStringFormat(404, "404 NOT FOUND: %s\n", ctx.Path)
		return
	}

	//将请求路由对应的处理器添加到上下文的中间件
	ctx.middlewares = append(ctx.middlewares, self.handlers[path])
	ctx.middleNum = len(ctx.middlewares)
	ctx.NextMiddlewares() //按顺序执行中间件
}
