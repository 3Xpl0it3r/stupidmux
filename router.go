package stupidmux

import (
	"github.com/3Xpl0it3r/stupidmux/context"
	"net/http"
	"strings"
)

type router struct {
	handlers map[string]context.StupidHandler
	roots map[string]*node
}

func newRouter()*router{
	return &router{
		handlers: make(map[string]context.StupidHandler, 0),
		roots:    make(map[string]*node, 0),
	}
}

func (r *router)addRouter(method, pattern string, handler context.StupidHandler){
	r.handlers[routerKey(method, pattern)] = handler
	if _,ok := r.roots[method];!ok{
		r.roots[method] = &node{}
	}
	r.roots[method].insertChild(parsePath(pattern), pattern)
}

func(r *router)getRouter(method,path string)(*node, context.Params){
	var params = make(context.Params, 0)
	if _,ok := r.roots[method];!ok {
		return nil, nil
	}
	reqPath := parsePath(path)
	n := r.roots[method].searchChild(reqPath)
	if n == nil{
		return nil, nil
	}
	for index,value := range parsePath(n.pattern){
		if value[0] == '*'{
			params[value[1:]] = strings.Join(reqPath[index:], "/")
		}else if value[0] == ':' {
			params[value[1:]] = reqPath[index]
		}
	}
	return n, params
}

func(r *router)handle(ctx *context.Context){
	ctx.Next()
	n,params := r.getRouter(ctx.Method, ctx.Path)
	if n != nil {
		if handle,ok := r.handlers[routerKey(ctx.Method, n.pattern)];ok {
			ctx.Params = params
			handle(ctx)
		}
	} else {
		ctx.String(http.StatusNotFound, "404 Not Found")
	}
}


func parsePath(path string)[]string{
	paths := make([]string, 0)
	parsedPath := strings.Split(path, "/")
	for _,v  := range parsedPath{
		if v != ""{
			paths = append(paths, v)
		}
	}
	return paths
}

func routerKey(method,pattern string)string{
	return method + "-" + pattern
}

