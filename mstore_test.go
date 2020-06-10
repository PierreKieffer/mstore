package mstore

import (
	"reflect"
	"testing"
)

func TestInitStorage(t *testing.T) {
	s := InitStorage()
	sType := reflect.TypeOf(s).String()
	must := "*[]*mstore.Index"
	if sType != must {
		t.Errorf("Storage test type failed")
	}
}

func TestInsert(t *testing.T) {
	s := InitStorage()

	testDocument := Document{
		Key:  "key_ab",
		Data: "data_test_1",
	}

	testDocument2 := Document{
		Key:  "key_ba",
		Data: "data_test_2",
	}

	var insertedKey = false
	var insertedData = false

	var embInsertedKey = false
	var embInsertedData = false

	Insert(s, testDocument)
	Insert(s, testDocument2)

	for _, v := range *s {
		if v.Head != nil {
			if v.Head.NodeDocument.Key == testDocument2.Key {
				insertedKey = true
			}
			if v.Head.NodeDocument.Data == testDocument2.Data {
				insertedData = true
			}

			if v.Head.NextNode != nil {
				if v.Head.NextNode.NodeDocument.Key == testDocument.Key {
					embInsertedKey = true
				}
				if v.Head.NextNode.NodeDocument.Data == testDocument.Data {
					embInsertedData = true
				}
			}

		}
	}

	if insertedKey == false || insertedData == false || embInsertedKey == false || embInsertedData == false {
		t.Errorf("Insert test failed")
	}

}

func TestFind(t *testing.T) {
	s := InitStorage()

	testDocument := Document{
		Key:  "key_ab",
		Data: "data_test",
	}

	Insert(s, testDocument)

	doc, _ := Find(s, "key_ab")

	if doc != testDocument {
		t.Errorf("Find test failed")
	}

}

func TestUpdate(t *testing.T) {
	s := InitStorage()

	testDocument := Document{
		Key:  "key_ab",
		Data: "data_test",
	}

	Insert(s, testDocument)

	updateDocument := Document{
		Key:  "key_ab",
		Data: "data_test_update",
	}

	Update(s, updateDocument)

	doc, _ := Find(s, "key_ab")

	if doc != updateDocument {
		t.Errorf("Update test failed")
	}

}

func TestDelete(t *testing.T) {
	s := InitStorage()

	testDocument := Document{
		Key:  "key_ab",
		Data: "data_test_1",
	}

	var insertedKey = false
	var insertedData = false

	Insert(s, testDocument)
	Delete(s, testDocument.Key)

	for _, v := range *s {
		if v.Head != nil {
			if v.Head.NodeDocument.Key == testDocument.Key {
				insertedKey = true
			}
			if v.Head.NodeDocument.Data == testDocument.Data {
				insertedData = true
			}

		}
	}

	if insertedKey == true || insertedData == true {
		t.Errorf("Delete test failed")
	}

}
