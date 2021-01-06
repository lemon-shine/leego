package leego

import (
	"testing"
)

func TestParsePath(t *testing.T) {
	t.Log(parsePath("/"))
	t.Log(parsePath("/a"))
	t.Log(parsePath("/a/b"))
	t.Log(parsePath("/a/b/:c"))
	t.Log(parsePath("/a/b/:c/d"))
	t.Log(parsePath("/a/b/:c/d/:e"))
	t.Log(parsePath("/a/b/:c/d/*asdf"))
	t.Log(parsePath("/a/b/:c/d/*sdfsdf/:e"))
}

func TestOneTire(t *testing.T) {
	tree := newTrie()
	tree.insert("/")
	tree.insert("/aaa/bbb")
	tree.insert("/aaa/bbb/ccc")
	t.Log(tree.search("/aaa"))
}

func TestTrie(t *testing.T) {
	tree := newTrie()

	tree.insert("/")
	tree.insert("/a")
	tree.insert("/a/b/*p")
	tree.insert("/a/b/d")
	tree.insert("/a/b/e/*k/f")
	tree.insert("/g/:h")

	t.Log(tree.search("/"))
	t.Log(tree.search("/a"))
	t.Log(tree.search("/a/b"))
	t.Log(tree.search("/a/b/123"))
	t.Log(tree.search("/a/b/345/d"))
	t.Log(tree.search("/a/b/e/f"))
	t.Log(tree.search("/g/333"))
	t.Log(tree.search("/f/f/f"))
}
