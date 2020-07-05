// Write a Bubble Sort program in Go.
package main

import "fmt"

/*
BubbleSort - As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted order.
*/
func BubbleSort(input []int) {
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-1-i; j++ {
			Swap(input, j)
		}
	}
}

/*
Swap - A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice.
You should write a Swap() function which performs this operation.
Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
*/
func Swap(slice []int, index int) {
	if slice[index] > slice[index+1] {
		temp := slice[index+1]
		slice[index+1] = slice[index]
		slice[index] = temp
	}
}

func main() {
	// The program should prompt the user to type in a sequence of up to 10 integers.
	input := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Print("Enter integer ", i+1, ": ")
		fmt.Scan(&input[i])
	}
	fmt.Println("BEFORE: ", input)

	// The program should print the integers out on one line, in sorted order, from least to greatest.
	BubbleSort(input)
	fmt.Println("AFTER:  ", input)
}
