/*******************************************************************************
模块：上下文
作者：Lemine
时间：2020/03/07
*******************************************************************************/
package leego

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//定义简单的JSON数据格式
type J map[string]interface{}

//定义请求响应上下文
type Context struct {
	//原始上下文
	Response http.ResponseWriter
	Request  *http.Request

	//请求信息
	Path   string
	Method string

	//响应信息
	StatusCode int
}

func NewContext(resp http.ResponseWriter, req *http.Request) *Context {
	ctx := &Context{
		Response: resp,
		Request:  req,
		Path:     req.URL.Path,
		Method:   req.Method,
	}

	ctx.init()
	return ctx
}

func (self *Context) init() {
	self.Request.ParseForm()
}

//Query：查询URL参数
func (self *Context) Query(key string) string {
	return self.Request.URL.Query().Get(key)
}

//GetPostForm：获取POST表单
func (self *Context) GetPostForm(key string) string {
	return self.Request.PostFormValue(key)
}

//SetStatus：设置响应状态码
func (self *Context) SetStatus(code int) {
	self.StatusCode = code
	self.Response.WriteHeader(code)
}

//SetHeader：设置响应头
func (self *Context) SetHeader(key string, value string) {
	self.Response.Header().Set(key, value)
}

//ResponseFail：返回错误的请求信息
func (self *Context) ResponseFail(code int, err string) {
	self.SetStatus(code)
	self.Response.Write([]byte(err))
}

//ResponseJson：响应JSON数据
func (self *Context) ResponseJson(code int, jsonObj interface{}) {
	self.SetHeader("Content-Type", "text/json")
	self.SetStatus(code)

	encoder := json.NewEncoder(self.Response)
	if err := encoder.Encode(jsonObj); err != nil {
		http.Error(self.Response, err.Error(), 500)
	}
}

//ResponseBytes：响应byte数据
func (self *Context) ResponseBytes(code int, data []byte) {
	self.SetStatus(code)
	self.Response.Write(data)
}

//ResponseString：响应字符串数据
func (self *Context) ResponseString(code int, value string) {
	self.SetHeader("Content-Type", "text/plain")
	self.SetStatus(code)
	self.Response.Write([]byte(value))
}

//ResponseFormatString：响应格式化字符串数据
func (self *Context) ResponseFormatString(code int, format string, values ...interface{}) {
	self.SetHeader("Content-Type", "text/plain")
	self.SetStatus(code)
	self.Response.Write([]byte(fmt.Sprintf(format, values...)))
}
