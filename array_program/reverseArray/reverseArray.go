package reversearray

//{1, 2, 3, 4, 5}
func ReverseArray(arr []int) []int {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return arr
}
