package context

import (
	"fmt"
	"net/http"
)

type StupidHandler func(ctx *Context)

type Params map[string]string


type Context struct {
	Writer http.ResponseWriter
	Request *http.Request

	Method string
	Path string

	Params     map[string]string
	statusCode int

	// middlewares
	MiddleWares []StupidHandler
	index int
}

func NewContext(w http.ResponseWriter,r *http.Request)*Context {
	return &Context{
		Writer: w,
		Request: r,
		Path: r.URL.Path,
		Method: r.Method,
		index: -1,
		MiddleWares: make([]StupidHandler,0),
	}
}

func(c *Context)Next(){
	c.index++
	s := len(c.MiddleWares)
	for ;c.index <s ;c.index++ {
		c.MiddleWares[c.index](c)
	}
}

func(c *Context)SetStatusCode(code int){
	c.statusCode = code
	c.Writer.WriteHeader(code)
}

func(c *Context)ParamByKey(key string)string{
	return c.Params[key]
}


func(c *Context)QueryByKey(key string)string{
	return c.Request.URL.Query().Get(key)
}

func(c *Context)SetHeader(key,value string){
	c.Writer.Header().Set(key, value)
}

func(c *Context)String(code int,format string,values ...interface{}){
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatusCode(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}
