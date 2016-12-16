package gofp

// FilterFunc is function that takes interface{} and produce bool to decide whether an element of an array should be outputed
type FilterFunc func(xi interface{}) bool

/*Filter is afunction that takes FilterFunc and produce filtered array
	Example:
	// Filter Func
	var filter = func(xi interface{}) bool {
		return (xi.(int)%2 == 1)
	}
	// You need to convert it array of interface first
	arr1 := [...]int{1, 2, 3, 4, 5}
	arr2 := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		arr2[i] = v
	}
	result := Filter(filter, arr2)
**/
func Filter(filter FilterFunc, x []interface{}) []interface{} {
	result := make([]interface{}, 0, len(x))
	for _, v := range x {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}
