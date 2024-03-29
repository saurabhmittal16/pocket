package core

type key struct {
	title string
}

type value struct {
	data any
}

type Node struct {
	id       int
	elements map[key]value
}

func (node Node) Get(keyLabel string) any {
	k := key{keyLabel}
	element := node.elements[k]
	v := element.data
	return v
}
func (node Node) Put(keyLabel string, data any) {
	k := key{keyLabel}
	v := value{data}

	node.elements[k] = v
}

func CreateNode(id int) *Node {
	node := Node{id, map[key]value{}}
	return &node
}
