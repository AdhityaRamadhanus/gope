package gofp

// FilterFunc is function that takes interface{} and produce bool to decide whether an element of an array should be outputed
type FilterFunc func(xi interface{}) bool

// Filter is afunction that takes FilterFunc and produce filtered array
func Filter(filter FilterFunc, x []interface{}) []interface{} {
	result := make([]interface{}, 0, len(x))
	for _, v := range x {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}
