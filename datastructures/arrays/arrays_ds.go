package arrays

// Complete the reverseArray function below.
func reverseArray(a []int32) []int32 {
	var result []int32
	size := len(a)

	for i := 1; i <= size; i++ {
		result = append(result, a[size-i])
	}
	return result
}
