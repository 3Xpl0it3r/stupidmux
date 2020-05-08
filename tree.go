package stupidmux

import "strings"

type node struct {
	path string
	pattern string
	children []*node
	wildChild bool
}

func(n *node)insertChild(paths []string, pattern string){
	if len(paths) == 0{
		n.pattern = pattern
		return
	}
	child := n.matchChild(paths[0])
	if child == nil{
		child = &node{path: paths[0], wildChild: paths[0][0] == ':' || paths[0][0] == '*'}
		n.children = append(n.children, child)
	}
	child.insertChild(paths[1:], pattern)
}

func(n *node)searchChild(paths []string)*node{
	if len(paths) == 0 || strings.HasPrefix(n.path, ":") || strings.HasPrefix(n.path, "*") {
		if n.pattern != ""{
			return n
		}else {
			return nil
		}
	}
	children := n.matchChildren(paths[0])
	for _,child := range children{
		result := child.searchChild(paths[1:])
		if result != nil{
			return result
		}
	}
	return nil
}

func(n *node)matchChild(path string)*node{
	for _,child := range n.children{
		if child.path == path {
			return child
		}
	}
	return nil
}
func(n *node)matchChildren(path string)[]*node{
	nodes := make([]*node, 0)
	for _,child := range n.children{
		if child.path == path || child.wildChild{
			nodes = append(nodes, child)
		}
	}
	return nodes
}