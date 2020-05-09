package stupidmux

import (
	"fmt"
	"github.com/3Xpl0it3r/stupidmux/middeware"
	"github.com/stretchr/testify/suite"
	"github.com/3Xpl0it3r/stupidmux/context"
	"net/http"
	"net/http/httptest"
	"testing"
)



type TestStupidMuxSuite struct {
	suite.Suite
}


func(suite *TestStupidMuxSuite)TestStupidMux1(){
	mux := NewStupidMux()
	v1 := mux.Group("/v1")
	v2 := mux.Group("/v2")

	v1.Use(middeware.LogMiddleWare)

	v1.GET("/", func(ctx *context.Context) {
		ctx.String(http.StatusOK, "HelloWorld")
	})
	v1.GET("/foo", func(ctx *context.Context) {
		ctx.String(http.StatusOK, "index")
	})

	v1.GET("/foo/:name", func(ctx *context.Context) {
		ctx.String(http.StatusOK, fmt.Sprintf("name:%s", ctx.ParamByKey("name")))
	})
	v2.GET("/static", func(ctx *context.Context) {
		ctx.String(http.StatusOK, "static")
	})
	v2.GET("/static/:type", func(ctx *context.Context) {
		ctx.String(http.StatusOK, fmt.Sprintf("type:%s", ctx.ParamByKey("type")))
	})
	reqUrl := map[string]string{"/v1/":"HelloWorld", "/v1/foo": "index", "/v1/foo/java": "name:java",
		"/v2/static/javascript": "type:javascript"}
	for k,v := range reqUrl{
		req,err := http.NewRequest("GET", k, nil)
		suite.Assert().Nil(err)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)
		suite.Assert().Equal(rr.Code, http.StatusOK)
		suite.Assert().Equal(rr.Body.String(), v)
	}
}



func(suite *TestStupidMuxSuite)TestStupidMux2(){

	v1 := Group("/v1")
	v2 := Group("/v2")

	v1.GET("/", func(ctx *context.Context) {
		ctx.String(http.StatusOK, "HelloWorld")
	})
	v1.GET("/foo", func(ctx *context.Context) {
		ctx.String(http.StatusOK, "index")
	})

	v1.GET("/foo/:name", func(ctx *context.Context) {
		ctx.String(http.StatusOK, fmt.Sprintf("name:%s", ctx.ParamByKey("name")))
	})
	v2.GET("/static", func(ctx *context.Context) {
		ctx.String(http.StatusOK, "static")
	})
	v2.GET("/static/:type", func(ctx *context.Context) {
		ctx.String(http.StatusOK, fmt.Sprintf("type:%s", ctx.ParamByKey("type")))
	})
	reqUrl := map[string]string{"/v1/":"HelloWorld", "/v1/foo": "index", "/v1/foo/java": "name:java",
		"/v2/static/javascript": "type:javascript"}
	for k,v := range reqUrl{
		req,err := http.NewRequest("GET", k, nil)
		suite.Assert().Nil(err)
		rr := httptest.NewRecorder()
		stupidMux.ServeHTTP(rr, req)
		suite.Assert().Equal(rr.Code, http.StatusOK)
		suite.Assert().Equal(rr.Body.String(), v)
	}
}

func TestStupidMux(t *testing.T){
	suite.Run(t, new(TestStupidMuxSuite))
}
