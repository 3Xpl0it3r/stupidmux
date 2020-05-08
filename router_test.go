package stupidmux

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestRouterSuite struct {
	suite.Suite
}

func(suite *TestRouterSuite)TestRouter(){
	r := newRouter()
	var handler StupidHandler
	r.addRouter("get", "/", handler)
	r.addRouter("get", "/index", handler)
	r.addRouter("get", "/foo", handler)
	r.addRouter("get", "/foo/:name", handler)
	r.addRouter("get", "/static/*filetype", handler)

	var n *node
	var params Params
	n,params = r.getRouter("get", "/")
	suite.Assert().Equal(n.pattern, "/")

	n,params = r.getRouter("get", "/index")
	suite.Assert().Equal(n.pattern, "/index")

	n,params = r.getRouter("get", "/foo")
	suite.Assert().Equal(n.pattern, "/foo")

	n,params = r.getRouter("get", "/foo/kubernetes")
	suite.Assert().Equal(n.pattern, "/foo/:name")
	suite.Assert().Equal(params["name"], "kubernetes")

	n,params = r.getRouter("get", "/static/javascript")
	suite.Assert().Equal(n.pattern, "/static/*filetype")
	suite.Assert().Equal(params["filetype"], "javascript")
}

func TestRouter(t *testing.T){
	suite.Run(t, new(TestRouterSuite))
}
