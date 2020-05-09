package stupidmux

import (
	"github.com/3Xpl0it3r/stupidmux/context"
	"net/http"
	"strings"
	"sync"
)

var (
	stupidMux *StupidMux
	stupidOnce = &sync.Once{}
)

func init() {
	if stupidMux == nil{
		stupidOnce.Do(func() {
			stupidMux = NewStupidMux()
		})
	}
}

type StupidMux struct {
	*routerGroup
	router *router
}

func NewStupidMux()*StupidMux{
	mux := &StupidMux{
		router:      newRouter(),
	}
	mux.routerGroup = &routerGroup{stupidMux: mux, isRoot: true}
	return mux
}



func(mux *StupidMux)GET(pattern string, handler context.StupidHandler){
	mux.router.addRouter("GET", pattern, handler)
}

func(mux *StupidMux)POST(pattern string, handler context.StupidHandler){
	mux.router.addRouter("POST", pattern, handler)
}

func(mux *StupidMux)DELETE(pattern string, handler context.StupidHandler){
	mux.router.addRouter("DELETE", pattern, handler)
}

func(mux *StupidMux)Run(address string)error{
	return http.ListenAndServe(address, mux)
}

func(mux *StupidMux)ServeHTTP(w http.ResponseWriter,r *http.Request){
	c := context.NewContext(w, r)
	for _,group := range mux.groups{
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			c.MiddleWares = append(c.MiddleWares, group.middleWares...)
		}
	}
	mux.router.handle(c)
}



// for

func GET(pattern string, handler context.StupidHandler){
	stupidMux.router.addRouter("GET", pattern, handler)
}

func POST(pattern string, handler context.StupidHandler){
	stupidMux.router.addRouter("POST", pattern, handler)
}

func DELETE(pattern string, handler context.StupidHandler){
	stupidMux.router.addRouter("DELETE", pattern, handler)
}

func Run(addr string)error{
	return http.ListenAndServe(addr, stupidMux)
}