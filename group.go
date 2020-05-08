package stupidmux

import "github.com/3Xpl0it3r/stupidmux/context"

type routerGroup struct {
	//middlewares
	middleWares []context.StupidHandler
	//
	prefix string
	parent *routerGroup
	groups []*routerGroup
	// only record router
	stupidMux *StupidMux
	// is root group
	isRoot bool
}

func Group(prefix string)*routerGroup{
	newGroup := &routerGroup{
		prefix:      prefix,
		parent:      stupidMux.routerGroup,
		stupidMux:   stupidMux,
		isRoot:      true,
	}
	stupidMux.groups = append(stupidMux.groups, newGroup)
	return newGroup
}

func(group *routerGroup)Group(prefix string)*routerGroup{
	mux := group.stupidMux
	newGroup := &routerGroup{
		prefix:      prefix,
		parent:      group,
		stupidMux:         mux,
		isRoot: false,
	}
	group.groups = append(group.groups, newGroup)
	return newGroup
}

func(group *routerGroup)addRoute(method, comps string, handler context.StupidHandler){
	var path string = ""
	var curGroup = group
	for  {
		path = curGroup.prefix + path
		if curGroup.isRoot {
			break
		}else {
			 curGroup = curGroup.parent
		}
	}
	group.stupidMux.router.addRouter(method, path+ comps, handler)
}

func(group *routerGroup)GET(subPattern string, handler context.StupidHandler){
	group.addRoute("GET", subPattern, handler)
}

func(group *routerGroup)POST(subPattern string,handler context.StupidHandler){
	group.addRoute("POST", subPattern, handler)
}

func(group *routerGroup)DELETE(subPattern string, handler context.StupidHandler){
	group.addRoute("DELETE", subPattern, handler)
}

// middleWares
func(group *routerGroup)Use(middleWares ...context.StupidHandler){
	group.middleWares = append(group.middleWares, middleWares...)
}