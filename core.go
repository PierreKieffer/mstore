package mstore

func InitStorage(customSize ...int) *[]*Index {
	var size int
	if len(customSize) > 0 {
		size = customSize[0]
	} else {
		size = 1000
	}
	storage := make([]*Index, size)

	for index, _ := range storage {
		storage[index] = &Index{}
	}
	return &storage
}

type Index struct {
	Head *Node
}

type Node struct {
	NodeDocument Document
	NextNode     *Node
}

type Document struct {
	Key  string
	Data interface{}
}
