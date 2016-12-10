package gofp

// Mapper is function that map a value to another value
type Mapper func(xi interface{}) interface{}

// Map is a function that takes mapper and array of interface and produce new array
func Map(mapper Mapper, x []interface{}) []interface{} {
	result := make([]interface{}, len(x), len(x))
	for i, v := range x {
		result[i] = mapper(v)
	}
	return result
}
