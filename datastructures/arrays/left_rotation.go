package arrays

/*
 * Complete the 'rotateLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER d: rotate the array d times
 *  2. INTEGER_ARRAY arr
 */

func RotateLeft(d int32, arr []int32) []int32 {
	return append(arr[d:], arr[0:d]...)
}
