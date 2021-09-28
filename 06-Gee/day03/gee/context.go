package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// context 作用为帮助用户添加一些额外信息
// 例如： Header/StatusCode 等

type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Request *http.Request
	Method string
	Path string
	Params map[string]string
}

func NewContext(writer http.ResponseWriter, request *http.Request)  *Context {
	return &Context{
		Writer: writer,
		Request: request,
		Method: request.Method,
		Path: request.URL.Path,
	}
}

func (context *Context) PostForm(key string) string {
	return context.Request.FormValue(key)
}

func (context *Context) Param(key string) string {
	return context.Request.URL.Query().Get(key)
}

func (context *Context) Query(key string) string {
	return context.Request.URL.Query().Get(key)
}

func (context *Context) Status(code int)  {
	context.Writer.WriteHeader(code)
}

func (context *Context) SetHeader(key string, value string)  {
	context.Writer.Header().Set(key, value)
}

func (context *Context) String(code int, format string, values...interface{})  {
	context.SetHeader("Content-Type", "text/plain")
	context.Status(code)
	context.Writer.Write([]byte(fmt.Sprintf(format, values)))
}


func (context *Context) JSON(code int, obj interface{})  {
	context.SetHeader("Content-Type", "application/json")
	context.Status(code)

	encoder := json.NewEncoder(context.Writer)

	if err := encoder.Encode(obj); err != nil {
		http.Error(context.Writer, err.Error(), 500)
	}
}

func (context *Context) Data(code int, data []byte)  {
	context.Status(code)
	context.Writer.Write(data)
}

func (context *Context) HTML(code int, html string)  {
	context.Status(code)
	context.SetHeader("Content-Type", "text/html")
	context.Writer.Write([]byte(html))
}