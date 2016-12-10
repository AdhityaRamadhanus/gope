package gofp

type Reducer func(acc interface{}, xi interface{}) interface{}

func Reduce(r Reducer, x []interface{}, init interface{}) interface{} {
	var result = init
	for _, v := range x {
		result = r(result, v)
	}
	return result
}
