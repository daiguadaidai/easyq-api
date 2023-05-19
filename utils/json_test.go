package utils

import (
	"fmt"
	"testing"
)

type ObjJson struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func TestObjToMap(t *testing.T) {
	obj := &ObjJson{
		Id:   1,
		Name: "name1",
		Age:  20,
	}

	objMap, err := ObjToMap(obj)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(objMap)
}
