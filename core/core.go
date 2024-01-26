package core

import "github.com/google/uuid"

type key struct {
	title string
}

type value struct {
	value any
}

type Node struct {
	id    uuid.UUID
	elems map[key]value
}

func (node Node) Get(_key string) any {
	val := node.elems[key{_key}]
	return val.value
}

func (node Node) Put(_key string, _value any) {
	node.elems[key{_key}] = value{value: _value}
}

func CreateNode() *Node {
	node := Node{uuid.New(), map[key]value{}}
	return &node
}
