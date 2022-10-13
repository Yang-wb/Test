package main

import (
	"errors"
	"fmt"
	"net/http"
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
	handler, found := h.findRouter(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found"))
		return
	}
	handler(c)
}

func (h *HandlerBasedOnTree) findRouter(path string) (handlerFunc, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, path := range paths {
		mathChild, found := cur.findMatchChild(path)
		if !found {
			return nil, false
		}
		cur = mathChild
	}

	if cur.handler == nil {
		return nil, false
	}

	return cur.handler, true
}

func (h *HandlerBasedOnTree) Route(method string, pattern string, handlerFunc func(ctx *Context)) {

	err := h.validatePattern(pattern)
	if err != nil {
		fmt.Printf("error : %v", err)
	}

	// 去掉前后的 /
	pattern = strings.Trim(pattern, "/")
	// paths [user, friends]
	paths := strings.Split(pattern, "/")

	cur := h.root

	for index, path := range paths {
		mathChild, ok := cur.findMatchChild(path)
		if ok {
			cur = mathChild
		} else {
			h.createSubTree(cur, paths[index:], handlerFunc)
			return
		}
	}
}

func (n *node) findMatchChild(path string) (*node, bool) {
	var wildcardNode *node
	for _, child := range n.children {
		// 并不是 * 的节点命中了，直接返回
		// ！= * 是为了防止用户乱输入
		if child.path == path && child.path != "*" {
			return child, true
		}
		// 命中了通配符的，我们看看后面还有没有更加详细的
		if child.path == "*" {
			wildcardNode = child
		}
	}
	return wildcardNode, wildcardNode != nil
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

func (h *HandlerBasedOnTree) validatePattern(pattern string) error {
	// 校验 *,如果存在，必须是最后一个，并且它前面必须是/
	// 即我们只接受 /* 的存在，abc*这种是非法的

	pos := strings.Index(pattern, "*")
	// 找到了*
	if pos > 0 {
		// 必须是最后一个
		if pos != len(pattern)-1 {
			return errors.New("ErrorInvalidRouterPattern")
		}
		if pattern[pos-1] != '/' {
			return errors.New("ErrorInvalidRouterPattern")
		}
	}

	return nil
}

func NewHandlerBasedOnTree() Handler {
	return &HandlerBasedOnTree{
		root: newNode("/"),
	}
}
