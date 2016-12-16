package gofp

// Mapper is function that map a value to another value
type Mapper func(xi interface{}) interface{}

/*Map is a function that takes mapper and array of interface and produce new array
	Example:
	// Mapper func
	var mapper1 = func(xi interface{}) interface{} {
		return xi.(int) * 2
	}
	// You need to convert it array of interface first
	arr1 := [...]int{1, 2, 3, 4, 5}
	arr2 := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		arr2[i] = v
	}
	result := Map(mapper1, arr2)
**/
func Map(mapper Mapper, x []interface{}) []interface{} {
	result := make([]interface{}, len(x), len(x))
	for i, v := range x {
		result[i] = mapper(v)
	}
	return result
}
