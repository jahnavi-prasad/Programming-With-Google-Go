package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, e int
	fmt.Printf("How many numbers you want to enter (min=4) : ")

	fmt.Scanf("%d", &n)
	if n < 4 {
		panic("Min 4 number of ints required")
	}
	nums := make([]int, 0)
	fmt.Printf("Enter number and hit return/enter\n")

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &e)
		nums = append(nums, e)
	}
	fmt.Printf("After input => %v \n", nums)

	var res []int
	ch := make(chan []int)
	step := n / 4

	for i := 0; i < n; i += step {
		go sortNums(nums[i:i+step], ch)
	}

	for i := 0; i < 4; i++ {
		rs1 := <-ch
		res = merge(res, rs1)
	}

	sort.Ints(res)
	fmt.Printf("Final Sorted output %v\n", res)
}

func sortNums(nums []int, ch chan []int) {
	fmt.Printf("Input => %v \n", nums)
	sort.Ints(nums)
	fmt.Printf("Sorted => %v \n", nums)
	ch <- nums
}

func merge(left, right []int) []int {

	length, i, j := len(left)+len(right), 0, 0
	slice := make([]int, length, length)

	for k := 0; k < length; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
