package main

import (
	"array_program/program"
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Finding the maximum and minimum values in the array
	max, min := program.FindMaxandMin(arr)
	fmt.Printf("Max: %d, Min: %d\n", max, min)

	arr1 := []int{1, 2, 3, 4, 5}
	// Printing the original array
	fmt.Println("Original Array:", arr1)
	fmt.Println("Reversed Array:", program.ReverseArray(arr1))

	// Check if an Array is Sorted
	arr2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Is Sorted:", program.IsSorted(arr2))

	// Finding the second maximum value in the array
	arr3 := []int{1, 2, 3, 4, 5}
	fmt.Println("Second Largest:", program.SecondLarget(arr3))

	// Sorting an Array of Integers
	arr4 := []int{5, 2, 6, 3, 1, 4}

	sort.Ints(arr4)

	// Print the sorted array
	fmt.Println("Sorted array in ascending order:", arr4)

	sort.Sort(sort.Reverse(sort.IntSlice(arr4)))

	// Print the sorted array
	fmt.Println("Sorted array in descending order:", arr4)

	// Remove Duplicates from an Array
	arr5 := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println("Original array:", arr5)
	fmt.Println("Array after removing duplicates:", program.RemoveDuplicates(arr5))

	// Inserting a an element into an Slice
	slice := []int{1, 2, 4, 5}
	fmt.Println("Original slice:", slice)
	fmt.Println("Slice after inserting", program.InsertingNewElement(100, 2, slice))

	// Finding the index of an element in a Slice
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Index of element 3 in slice:", program.FindIndex(slice1, 3))

	//  Move All Zeros to the End of an Array

	arr6 := []int{0, 1, 0, 3, 12}
	fmt.Println("Original array:", arr6)
	fmt.Println("Array after moving zeros to the end:", program.MoveZerosToEnd(arr6))

	// Best Time to Buy and Sell Stock
	// Problem Statement: You are given an array prices where prices[i] is the price of a given
	// stock on the ith day. You want to maximize your profit by choosing a single day to buy one stock
	// and choosing a different day in the future to sell that stock. Return the maximum profit
	// you can achieve from this transaction. If you cannot achieve any profit, return 0.
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Best Time to Buy and Sell Stock:", program.MaxProfit(prices))

}
