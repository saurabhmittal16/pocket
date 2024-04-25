package core

import (
	"fmt"
	"testing"
)

func assert(expected, actual any) string {
	return fmt.Sprintf(`Expected: %v, Actual: %v`, expected, actual)
}

func TestGetAndPut(t *testing.T) {
	node := CreateStore("key1")
	node.Put("a", "1")
	node.Put("b", "2")

	val := node.Get("a")
	if val != "1" {
		t.Fatal("Node GET returns incorrect value. " + assert(1, val))
	}

	val = node.Get("x")
	if len(val) > 0 {
		t.Fatal("Node GET returns incorrect value. " + assert(nil, val))
	}
}
