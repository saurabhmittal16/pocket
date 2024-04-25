package core

type key struct {
	title string
}

type value struct {
	data string
}

type Store struct {
	id       string
	elements map[key]value
}

func (node Store) Get(keyLabel string) string {
	k := key{keyLabel}
	element := node.elements[k]
	v := element.data
	return v
}
func (node Store) Put(keyLabel string, data string) {
	k := key{keyLabel}
	v := value{data}

	node.elements[k] = v
}

func CreateStore(id string) *Store {
	store := Store{id, map[key]value{}}
	return &store
}
