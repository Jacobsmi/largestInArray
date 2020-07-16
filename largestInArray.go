package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getLargestItem(arr [5]int) int {
	// Create a variable to hold the largest item
	largest := arr[0]
	// Store the index of the largest item
	index := 0
	// For all items in the array search for a larger item while tracking the largest and the index of it
	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
			index = i
		}
	}
	// Return the index of the largest
	return index
}

func getRandomArray() [5]int {
	// Declaring an array that will hold the random numbers
	var arr [5]int
	// Seed the random object with a pseudo random seed
	rand.Seed(time.Now().UnixNano())
	// Populate the array with sudo-random numbers
	for i := 0; i < 5; i++ {
		arr[i] = rand.Intn(100)
	}
	// Return the pseudo random array
	return arr
}

func main() {
	// Populate a pseudo-random array
	arr := getRandomArray()
	fmt.Println("The array is", arr)
	// Find the index of the largest element from getLargestItem
	indexOfLargest := getLargestItem(arr)
	fmt.Println("The largest item is", arr[indexOfLargest], "at index", indexOfLargest)
}
