package tools

import (
	"fmt"
	"testing"
)

func TestStructToMap(t *testing.T) {
	type testA struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var a = testA{
		Name: "111",
		Age:  2,
	}
	dist := StructToMap(a)
	fmt.Println(dist)
}
