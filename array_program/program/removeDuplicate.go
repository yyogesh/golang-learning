package program

import (
	"fmt"
	"sort"
)

func RemoveDuplicates(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	sort.Ints(arr)
	fmt.Println("after sorting", arr)
	i := 1
	//{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	// {1, 1, 2, 3, 3, 4, 5, 5, 6, 9}
	for j := 1; j < len(arr); j++ {
		if arr[j] != arr[j-1] {
			arr[i] = arr[j]
			i++
		}
	}
	fmt.Println("before removing", arr)
	return arr[:i]
}
