package gofp

type reducer func(acc interface{}, xi interface{}) interface{}

func Reduce(r reducer, x []interface{}, init interface{}) interface{} {
	var result = init
	for _, v := range x {
		result = r(result, v)
	}
	return result
}
