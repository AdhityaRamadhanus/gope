package gofp

type ComposableFunc func(x interface{}) interface{}

func Compose(funs ...ComposableFunc) ComposableFunc {
	lenfun := len(funs)
	reversedfuns := make([]ComposableFunc, lenfun, lenfun)
	for i, v := range funs {
		reversedfuns[lenfun-i-1] = v
	}
	return func(x interface{}) interface{} {
		var result = x
		for _, v := range funs {
			result = v(result)
		}
		return result
	}
}
