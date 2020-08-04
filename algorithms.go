package main

import (
	"fmt"
	"math/rand"
	"time"
)

func swap(a *int, b *int) {
	var temp int

	temp = *a
	*a = *b
	*b = temp
}

func bubblesort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
	return arr
}

func linearsearchConcurrent(arr []int, searchNum int, channel chan int) {

	for i := 0; i < len(arr); i++ {
		if arr[i] == searchNum {
			channel <- 0
		}
	}
	channel <- 1
}

func main() {

	// to generate 'trully' random numbers and not biased ones
	rand.Seed(time.Now().Unix())

	// initializing a empty slice for the random numbers
	randNumber := []int{}

	// create a slice with positve numbers
	for i := 0; i < 100; i++ {
		randNumber = append(randNumber, rand.Intn(500))
	}

	fmt.Println("Sorted Array:", bubblesort(randNumber))

	// creating a slice to pass it as a parameter to the search function
	array := []int{3, 2, 1, 5, 6, 8}

	searchNumber := 5

	// create a channel
	channel := make(chan int)

	go linearsearchConcurrent(array[:len(array)/2], searchNumber, channel)
	go linearsearchConcurrent(array[len(array)/2:], searchNumber, channel)

	notFound := <-channel

	if notFound == 1 {
		fmt.Printf("Number %d was not found.", searchNumber)
	} else {
		fmt.Printf("Number %d was found.", searchNumber)
	}
}
