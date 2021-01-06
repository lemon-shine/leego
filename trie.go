/*******************************************************************************
Method: 路由树
Author: Lemine
Langua: Golang 1.14
Modify：2020/03/07
*******************************************************************************/
package leego

import (
	"strings"
)

type node struct {
	//路由路径
	path string
	//路由"/"分割的部分
	part string
	//是否动态匹配，":"和"*"开头的部分为动态匹配
	isWild   bool
	children map[string]*node
}

type trie struct {
	root *node
}

func newTrie() *trie {
	return &trie{
		root: &node{children: make(map[string]*node)},
	}
}

//插入新的路由节点
func (self *trie) insert(path string) {
	parts := parsePath(path)

	child := self.root

	for _, part := range parts {
		//判断part节点是否存在
		n, ok := child.children[part]
		if !ok {
			//添加新的part节点
			child.children[part] = &node{
				part:     part,
				isWild:   part[0] == ':' || part[0] == '*',
				children: make(map[string]*node),
			}

			child = child.children[part]
			continue
		}
		child = n
	}

	child.path = path
}

//搜索已注册路由节点
func (self *trie) search(path string) (*node, map[string]string) {
	parts := parsePath(path)
	params := make(map[string]string)

	child := self.root

Next:
	for _, part := range parts {
		//判断part是否精确参数
		if node, ok := child.children[part]; ok {
			child = node
			continue
		}

		//判断part是否为动态参数(遍历当前结点的所有子节点)
		for _, child = range child.children {
			if child.isWild {
				params[child.part[1:]] = part
				continue Next
			}
		}

		//part参数不存在
		return nil, nil
	}

	//判断检索到的路由节点是否注册
	//不能使用child.path==path，path可能是动态的
	if child.path == "" {
		return nil, nil
	}

	return child, params
}

//解析URL
func parsePath(path string) []string {
	results := make([]string, 0)
	parts := strings.Split(path, "/")

	for _, part := range parts {
		if part != "" {
			results = append(results, part)

			if part[0] == '*' {
				return results
			}
		}
	}

	return results
}
