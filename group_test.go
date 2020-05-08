package stupidmux

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestRouterGroupSuite struct {
	suite.Suite
}

func(suite *TestRouterGroupSuite)TestRouterGroup(){


}

func TestRouterGroup(t *testing.T){
	suite.Run(t, new(TestRouterSuite))
}
