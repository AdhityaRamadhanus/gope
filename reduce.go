package gope

// Reducer is a function that takes accumulated value and current value and produce result
type Reducer func(acc interface{}, xi interface{}) interface{}

/*Reduce takes a Reducer, array of interface and initial value to produce a single value from that array
	Example:
	// Reducer FUn
	var reducer1 = func(acc interface{}, xi interface{}) interface{} {
		return acc.(int) + xi.(int)
	}
	// You need to convert it slice of interface first
	arr1 := [...]int{1, 2, 3, 4, 5}
	arr2 := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		arr2[i] = v
	}
	result := Reduce(reducer1, arr2, 0)
**/
func Reduce(r Reducer, x []interface{}, init interface{}) interface{} {
	var result = init
	for _, v := range x {
		result = r(result, v)
	}
	return result
}
