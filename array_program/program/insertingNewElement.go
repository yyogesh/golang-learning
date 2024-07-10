package program

func InsertingNewElement(element, index int, arr []int) []int {
	// append([]int{element}, slice[index:]})
	// {1, 2, 100, 4, 5}
	return append(arr[:index], append([]int{element}, arr[index:]...)...)
}
