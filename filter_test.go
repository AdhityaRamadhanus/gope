package gope

import (
	"testing"
)

func TestFilter(t *testing.T) {
	var filter = func(xi interface{}) bool {
		return (xi.(int)%2 == 1)
	}
	arr1 := [...]int{1, 2, 3, 4, 5}
	arr2 := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		arr2[i] = v
	}
	result := Filter(filter, arr2)
	if len(result) != 3 {
		t.Error("Mapped array does not produce equal length ")
	}
	for _, v := range result {
		if v.(int)%2 != 1 {
			t.Error("Mapper doesn't produce expected value")
		}
	}
}
