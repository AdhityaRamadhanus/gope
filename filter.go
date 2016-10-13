package gofp

func Filter(filter func(xi interface{}) bool, x []interface{}) []interface{} {
	result := make([]interface{}, 0, len(x))
	for _, v := range x {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}
