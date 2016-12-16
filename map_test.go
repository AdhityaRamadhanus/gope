package gope

import (
	"testing"
)

func TestMap(t *testing.T) {
	var mapper1 = func(xi interface{}) interface{} {
		return xi.(int) * 2
	}
	arr1 := [...]int{1, 2, 3, 4, 5}
	arr2 := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		arr2[i] = v
	}
	result := Map(mapper1, arr2)
	if len(result) != len(arr1) {
		t.Error("Mapped array does not produce equal length ")
	}
	for i, v := range result {
		if v != arr1[i]*2 {
			t.Error("Mapper doesn't produce expected value")
		}
	}
}
