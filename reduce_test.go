package gofp

import (
	"testing"
)

func TestReduce(t *testing.T) {
	var reducer1 = func(acc interface{}, xi interface{}) interface{} {
		return acc.(int) + xi.(int)
	}
	arr1 := [...]int{1, 2, 3, 4, 5}
	arr2 := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		arr2[i] = v
	}
	result := Reduce(reducer1, arr2, 0)
	if result != 15 {
		t.Error("Reduced array should return 15")
	}
}
