package gee

import "strings"

// 前缀树
// 内部使用 也无须暴露出来
// /:lang/doc  可以匹配 /golang/doc, /c/doc 等
// /about
// /p/blog
// static/*filePath 可匹配 static/fav.ico 与 static/swift/hello.swift

// 例如： /:lang/doc
// 		/:lang/tutorial
//	    /:lang/intro
//		/about
//		/p/blog
//		/p/related
// 结构图如下
//							/
//		/:lang				/about 					/p
//  /intro /tutorial /doc 					 /blog	 	/releated

type node struct {
	pattern string
	part string
	children []*node
	isWild bool
}

func (node *node) matchChild(part string) *node {
	for _, child := range node.children {
		if child.part == part {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int)  {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)

	if child == nil {
		child = &node{ part: part,  isWild: part[0] == ':' || part[0] == '*' }
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return  nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

