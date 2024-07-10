package program

import "fmt"

func MoveZerosToEnd(arr []int) []int {
	n := len(arr)
	j := 0

	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			arr[j] = arr[i]
			j++
		}
	}

	fmt.Println("arr: ", arr, j)

	for j < n {
		arr[j] = 0
		j++
	}

	return arr
}
