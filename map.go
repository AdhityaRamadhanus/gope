package gofp

type Mapper func(xi interface{}) interface{}

func Map(mapper Mapper, x []interface{}) []interface{} {
	result := make([]interface{}, len(x), len(x))
	for i, v := range x {
		result[i] = mapper(v)
	}
	return result
}
