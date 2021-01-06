/*******************************************************************************
Method: 上下文
Author: Lemine
Langua: Golang 1.14
Modify：2020/03/07
*******************************************************************************/
package leego

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

//定义JSON数据类型
type JSON map[string]interface{}

//定义上下文的数据类型
type Context struct {
	engine *Engine

	//原始上下文
	Response http.ResponseWriter
	Request  *http.Request

	//请求信息
	Path   string
	Method string
	Params map[string]string

	//响应信息
	StatusCode int

	//中间件
	middlewares []HandleFunc //中间件
	middleNum   int          //中间件数量
	index       int          //中间索引，记录执行到第几个中间件
}

func newContext(engine *Engine, req *http.Request, resp http.ResponseWriter) *Context {
	return &Context{
		engine:   engine,
		Request:  req,
		Response: resp,
		Path:     req.URL.Path,
		Method:   req.Method,
		index:    -1,
	}
}

//获取URL的查询参数
func (self *Context) Query(key string) string {
	return self.Request.URL.Query().Get(key)
}

//获取URL的动态路由参数
func (self *Context) GetParam(key string) string {
	value, _ := self.Params[key]
	return value
}

//获取POST表单参数
func (self *Context) GetPostForm(key string) (string, error) {
	//检查表单是否已解析
	if self.Request.PostForm == nil {
		if err := self.Request.ParseForm(); err != nil {
			return "", err
		}
	}
	return self.Request.PostFormValue(key), nil
}

//获取所有的请求cookie
func (self *Context) GetCookies() []*http.Cookie {
	return self.Request.Cookies()
}

//获取指定的请求cookie，若有多个，则只返回第一个匹配cookie
func (self *Context) GetCookie(name string) (string, error) {
	cookie, err := self.Request.Cookie(name)
	if err != nil {
		return "", err
	}
	//将QueryEscape转码的字符串安全还原
	value, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return "", err
	}
	return value, nil
}

//获取多部分表单中匹配的第一个文件的头信息
func (self *Context) GetFormFileHeader(name string) (*multipart.FileHeader, error) {
	//检查多部分表单是否已解析
	if self.Request.MultipartForm == nil {
		//解析MultipartForm
		err := self.Request.ParseMultipartForm(self.engine.maxBodyBytes)
		if err != nil {
			return nil, err
		}
	}
	//获取多部分表单解析出的第一个目标文件及其头信息
	file, header, err := self.Request.FormFile(name)
	if err != nil {
		return nil, err
	}
	file.Close()
	return header, err
}

//获取已解析多部分表单，包括文件上传
func (self *Context) GetMultipartForm() (*multipart.Form, error) {
	if self.Request.MultipartForm == nil {
		//解析MultipartForm
		err := self.Request.ParseMultipartForm(self.engine.maxBodyBytes)
		if err != nil {
			return nil, err
		}
	}
	return self.Request.MultipartForm, nil
}

//------------------------------------------------------------------------------

//设置响应状态码
func (self *Context) SetStatus(code int) {
	self.StatusCode = code
	self.Response.WriteHeader(code)
}

//设置响应头
func (self *Context) SetHeader(key string, value string) {
	self.Response.Header().Set(key, value)
}

//设置Cookie
func (self *Context) SetCookie(name, value, path, domain string, maxAge int,
	secure, httpOnly bool, sameSite http.SameSite, raw string, unparsed []string) {
	if path == "" {
		path = "/"
	}
	http.SetCookie(self.Response, &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		Secure:   secure,
		HttpOnly: httpOnly,
		SameSite: sameSite,
		Raw:      raw,
		Unparsed: unparsed,
	})
}

//保存上传文件到指定路径
func (self *Context) SaveFormFile(file *multipart.FileHeader, dst string) error {
	//打开上传文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//创建本地文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	//将上传的文件内容拷贝到新建的本地文件中
	_, err = io.Copy(out, src)
	return err
}

//依次执行中间件
func (self *Context) NextMiddlewares() {
	self.index++
	for ; self.index < self.middleNum; self.index++ {
		self.middlewares[self.index](self)
	}
}

//------------------------------------------------------------------------------

//响应byte数据
func (self *Context) ResponseBytes(code int, data []byte) {
	self.SetStatus(code)
	self.Response.Write(data)
}

//响应JSON数据
func (self *Context) ResponseJSON(code int, jsonObj interface{}) {
	self.SetHeader("Content-Type", "text/json")
	self.SetStatus(code)

	encoder := json.NewEncoder(self.Response)
	if err := encoder.Encode(jsonObj); err != nil {
		http.Error(self.Response, err.Error(), 500)
	}
}

//响应字符串数据
func (self *Context) ResponseString(code int, value string) {
	self.SetHeader("Content-Type", "text/plain")
	self.SetStatus(code)
	self.Response.Write([]byte(value))
}

//响应格式化字符串数据
func (self *Context) ResponseStringFormat(code int, format string, values ...interface{}) {
	self.SetHeader("Content-Type", "text/plain")
	self.SetStatus(code)
	self.Response.Write([]byte(fmt.Sprintf(format, values...)))
}

func (self *Context) ResponseHTML(code int, name string, data interface{}) {
	self.SetHeader("Content-Type", "text/html")
	self.SetStatus(code)
	if err := self.engine.htmlTemplates.ExecuteTemplate(self.Response, name, data); err != nil {
		self.ResponseString(500, err.Error())
	}
}

// 将文件直接写入响应体
func (self *Context) ResponseFile(filepath string) {
	http.ServeFile(self.Response, self.Request, filepath)
}

//将文件写入响应体，在客户端将以给定名下载该文件
func (self *Context) ResponseFileAttachment(filepath, filename string) {
	self.SetHeader("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	http.ServeFile(self.Response, self.Request, filepath)
}

//将文件从http.FileSystem写入到响应体
func (self *Context) ResponseFileFromFileSystem(filepath string, fs http.FileSystem) {
	defer func(old string) {
		self.Request.URL.Path = old
	}(self.Request.URL.Path)

	self.Request.URL.Path = filepath
	http.FileServer(fs).ServeHTTP(self.Response, self.Request)
}
