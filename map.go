package gofp

func Map(mapper func(xi interface{}) interface{}, x []interface{}) []interface{} {
	result := make([]interface{}, len(x), len(x))
	for i, v := range x {
		result[i] = mapper(v)
	}
	return result
}
