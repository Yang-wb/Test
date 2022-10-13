package main

import (
	"strings"
)

type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node

	//如果这是叶子节点，
	//那么匹配上之后就可以调用该方法
	handler handlerFunc
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	panic("implement me")
}

func (h *HandlerBasedOnTree) Route(method string, pattern string, handlerFunc func(ctx *Context)) {
	// 去掉前后的 /
	pattern = strings.Trim(pattern, "/")
	// paths [user, friends]
	paths := strings.Split(pattern, "/")

	cur := h.root

	for index, path := range paths {
		mathChild, ok := cur.findMMatchChild(path)
		if ok {
			cur = mathChild
		} else {
			h.createSubTree(cur, paths[index:], handlerFunc)
			return
		}
	}
}

func (n *node) findMMatchChild(path string) (*node, bool) {
	for _, child := range n.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}

func (h *HandlerBasedOnTree) createSubTree(root *node, paths []string, handlerFunc handlerFunc) {
	cur := root
	for _, path := range paths {
		nn := newNode(path)
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handlerFunc
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 8),
		handler:  nil,
	}
}

func NewHandlerBasedOnTree() Handler {
	return &HandlerBasedOnTree{
		root: newNode("/"),
	}
}
