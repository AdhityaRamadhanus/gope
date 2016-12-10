package gofp

// Reducer is a function that takes accumulated value and current value and produce result
type Reducer func(acc interface{}, xi interface{}) interface{}

// Reduce takes a Reducer, array of interface and initial value to produce a single value from that array
func Reduce(r Reducer, x []interface{}, init interface{}) interface{} {
	var result = init
	for _, v := range x {
		result = r(result, v)
	}
	return result
}
