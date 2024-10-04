package companyobserver

import (
	"fmt"
	"reflect"
	"testing"
)

type Value struct {
	Name   string
	Gender *string
}

func TestDetach(t *testing.T) {
	v1 := Value{Name: "Tom", Gender: new(string)}
	v2 := Value{Name: "Tom", Gender: new(string)}
	//  if v1 == v2 {
	if reflect.DeepEqual(v1, v2) {
		fmt.Println("相等")
		return
	}

	fmt.Println("不相等")
}
