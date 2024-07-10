package program

func SecondLarget(arr []int) int {
	largest := arr[0]       // 1
	secondLargest := arr[0] // 1
	// { 5, 2, 3, 7}
	// {1, 2, 3, 4, 5}
	for _, value := range arr {
		// value = 3
		// secondLargest = 4
		// largest = 5
		// 3 > 2
		if value > largest {
			secondLargest = largest
			largest = value
			// 1 > 1
		} else if value > secondLargest && value != largest {
			secondLargest = value
		}
	}
	return secondLargest
}
