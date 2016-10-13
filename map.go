package gofp

import (
	"fmt"
)

func Map(mapper func(xi interface{}) interface{}, x []interface{}) []interface{} {
	result := make([]interface{}, len(x), len(x))
	for i := 0; i < len(x); i++ {
		// fmt.Println(mapper(x[i]))
		result[i] = mapper(x[i])
	}
	return result
}
