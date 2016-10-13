package gofp

func Reduce(reducer func(acc interface{}, xi interface{}) interface{}, x []interface{}, init interface{}) interface{} {
	var result = init
	for _, v := range x {
		result = reducer(result, v)
	}
	return result
}
