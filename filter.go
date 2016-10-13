package gofp

import (
	"fmt"
)

func Filter(filter func(xi interface{}) bool, x []interface{}) []interface{} {
	result := make([]interface{}, 0, len(x))
	for i := 0; i < len(x); i++ {
		// fmt.Println(filter(x[i]))
		if filter(x[i]) {
			result = append(result, x[i])
		}
	}
	return result
}
