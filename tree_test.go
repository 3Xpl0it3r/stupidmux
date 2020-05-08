package stupidmux

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestTreeSuite struct {
	suite.Suite
}

func(suite *TestTreeSuite)TestTree(){
	head := &node{}
	head.insertChild([]string{}, "/")
	head.insertChild([]string{"index"}, "/index")
	head.insertChild([]string{"foo"}, "/foo")
	head.insertChild([]string{"foo", ":name"}, "/foo/:name")
	head.insertChild([]string{"static", "*language"}, "/static/*language")

	suite.Assert().Equal(head.searchChild([]string{}).pattern, "/")
	suite.Assert().Equal(head.searchChild([]string{"index"}).pattern, "/index")
	suite.Assert().Equal(head.searchChild([]string{"foo"}).pattern, "/foo")
	suite.Assert().Equal(head.searchChild([]string{"foo", "kubernetes"}).pattern, "/foo/:name")
	suite.Assert().Equal(head.searchChild([]string{"foo", "docker"}).pattern, "/foo/:name")
	suite.Assert().Equal(head.searchChild([]string{"static", "go"}).pattern, "/static/*language")
	suite.Assert().Equal(head.searchChild([]string{"static", "java"}).pattern, "/static/*language")

}

func TestTree(t *testing.T){
	suite.Run(t, new(TestTreeSuite))
}