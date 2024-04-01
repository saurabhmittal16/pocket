package core

type key struct {
	title string
}

type value struct {
	data any
}

type Store struct {
	id       int
	elements map[key]value
}

func (node Store) Get(keyLabel string) any {
	k := key{keyLabel}
	element := node.elements[k]
	v := element.data
	return v
}
func (node Store) Put(keyLabel string, data any) {
	k := key{keyLabel}
	v := value{data}

	node.elements[k] = v
}

func CreateStore(id int) *Store {
	store := Store{id, map[key]value{}}
	return &store
}
