package gofp

type filterFunc func(xi interface{}) bool

func Filter(filter filterFunc, x []interface{}) []interface{} {
	result := make([]interface{}, 0, len(x))
	for _, v := range x {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}
