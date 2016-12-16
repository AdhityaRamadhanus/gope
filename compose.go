package gope

// ComposableFunc is composable function that takes an interface{} and produce interface{}
type ComposableFunc func(x interface{}) interface{}

/*Compose is a function that return a ComposableFunction and takes array of ComposableFUnc
   	Right now it's only accept unary function
	Example:
	// Composable Func
	var cf1 = func(x interface{}) interface{} {
		return x.(int) * 5
	}
	// Composable Func
	var cf2 = func(x interface{}) interface{} {
		return x.(int) * 6
	}

	var composed = Compose(cf2, cf1)
	result := composed(5)
**/
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
