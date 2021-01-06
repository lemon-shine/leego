package leego

import (
	"testing"
)

func TestGroup(t *testing.T) {
	router := NewEngine()
	router.NewGroup("/aaa")
	router.ListenAndServe(":8080")
}
