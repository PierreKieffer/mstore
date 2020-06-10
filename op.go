package mstore

import (
	"errors"
	"strconv"
	"time"
)

/*
Find
*/
func Find(storage *[]*Index, targetKey string) (Document, error) {
	index := HashFunction(len(*storage), targetKey)
	current := (*storage)[index].Head
	if current == nil {
		return Document{}, errors.New("No index available for targeted key : " + targetKey)
	}

	if current.NodeDocument.Key == targetKey {
		return current.NodeDocument, nil
	} else {
		for current != nil {
			if current.NodeDocument.Key == targetKey {
				return current.NodeDocument, nil
				break
			} else {
				current = current.NextNode
			}
		}
	}
	return Document{}, errors.New("No document available for targeted key : " + targetKey)

}

/*
Insert
*/
func Insert(storage *[]*Index, insertDocument Document) error {
	if insertDocument.Key == "" {
		insertDocument.Key = GenerateKey()
		Insert(storage, insertDocument)
	} else {
		checkDuplicate := CheckDuplicateKey(storage, insertDocument.Key)
		if checkDuplicate {
			return errors.New("Duplicate key error : " + insertDocument.Key)
		}

		index := HashFunction(len(*storage), insertDocument.Key)

		newNode := &Node{NodeDocument: insertDocument}

		if (*storage)[index].Head == nil {
			(*storage)[index].Head = newNode
		} else {
			newNode.NextNode = (*storage)[index].Head
			(*storage)[index].Head = newNode
		}

	}
	return nil
}

/*
Update
*/
func Update(storage *[]*Index, updateDocument Document) error {
	if updateDocument.Key == "" {
		return errors.New("Empty keys are not allowed")
	}
	index := HashFunction(len(*storage), updateDocument.Key)
	current := (*storage)[index].Head
	if current == nil {
		return errors.New("No index available for targeted key : " + updateDocument.Key)
	}
	if current.NodeDocument.Key == updateDocument.Key {
		current.NodeDocument.Data = updateDocument.Data
		return nil
	} else {
		for current != nil {
			if current.NodeDocument.Key == updateDocument.Key {
				current.NodeDocument.Data = updateDocument.Data
				return nil
				break
			} else {
				current = current.NextNode
			}
		}
	}
	return errors.New("No document available for targeted key : " + updateDocument.Key)
}

/*
Delete
*/
func Delete(storage *[]*Index, targetKey string) error {
	index := HashFunction(len(*storage), targetKey)

	var deleted = false

	current := (*storage)[index]
	if current.Head == nil {
		return errors.New("No index available for targeted key : " + targetKey)
	}

	if current.Head.NodeDocument.Key == targetKey {
		deleted = true
		if current.Head.NextNode != nil {
			current.Head.NodeDocument = current.Head.NextNode.NodeDocument
			if current.Head.NextNode.NextNode != nil {
				current.Head.NextNode = current.Head.NextNode.NextNode
			} else {
				current.Head.NextNode = nil
			}
		} else {
			current.Head = nil
		}
	} else {
		currentHead := current.Head

		for currentHead != nil {
			if currentHead.NextNode != nil {
				if currentHead.NextNode.NodeDocument.Key == targetKey {
					deleted = true
					if currentHead.NextNode.NextNode != nil {

						currentHead.NextNode = currentHead.NextNode.NextNode
					} else {
						currentHead.NextNode = nil
					}
				}
			}
			currentHead = currentHead.NextNode
		}
	}
	if deleted {
		return nil

	} else {
		return errors.New("No document available for targeted key : " + targetKey)
	}
	return nil
}

/*
HashFunction
Used to create the index from the document key
based on the ascii value of the key and the size of the storage
*/
func HashFunction(storageSize int, key string) int {
	var index int

	for _, r := range key {
		index = index + int(r)
	}
	index = index % storageSize
	return index
}

/*
GenerateKey
Generate UnixNano string key to reference documents
*/
func GenerateKey() string {
	unixTime := time.Now().UnixNano()
	return strconv.FormatInt(unixTime, 10)
}

/*
CheckDuplicateKey
Used to check if the document key already exist for insert operation.
Several different keys can have the same storage index,
but it is not possible to have several documents with the same key
*/
func CheckDuplicateKey(storage *[]*Index, key string) bool {
	doc, _ := Find(storage, key)
	if doc.Key == key {
		return true
	} else {
		return false
	}
	return true
}
