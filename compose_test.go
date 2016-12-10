package gofp

import (
	"log"
	"testing"
)

func TestComposeOneFunction(t *testing.T) {
	var cf1 = func(x interface{}) interface{} {
		return x.(int) * 5
	}
	var composed = Compose(cf1, cf1)
	result := composed(5)
	if result != 125 {
		t.Error("Composed function should return 125")
	}
}

func TestComposeTwoFunction(t *testing.T) {
	var cf1 = func(x interface{}) interface{} {
		return x.(int) * 5
	}
	var cf2 = func(x interface{}) interface{} {
		return x.(int) * 6
	}
	var composed = Compose(cf2, cf1)
	result := composed(5)
	if result != 150 {
		t.Error("Composed function should return 150")
	}
}
