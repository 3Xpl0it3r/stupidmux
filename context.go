package stupidmux

import (
	"fmt"
	"net/http"
)

type StupidHandler func(ctx *Context)

type Params map[string]string

type Context struct {
	writer http.ResponseWriter
	request *http.Request

	method string
	path string

	params Params
	statusCode int

	// middlewares
	middlewares []StupidHandler
	index int
}

func newContext(w http.ResponseWriter,r *http.Request)*Context{
	return &Context{
		writer: w,
		request: r,
		path: r.URL.Path,
		method: r.Method,
		index: -1,
	}
}

func(c *Context)Next(){
	c.index++
	s := len(c.middlewares)
	for ;c.index <s ;c.index++ {
		c.middlewares[c.index](c)
	}
}

func(c *Context)SetStatusCode(code int){
	c.statusCode = code
	c.writer.WriteHeader(code)
}

func(c *Context)Params(key string)string{
	return c.params[key]
}


func(c *Context)Query(key string)string{
	return c.request.URL.Query().Get(key)
}

func(c *Context)SetHeader(key,value string){
	c.writer.Header().Set(key, value)
}

func(c *Context)String(code int,format string,values ...interface{}){
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatusCode(code)
	c.writer.Write([]byte(fmt.Sprintf(format, values...)))
}
