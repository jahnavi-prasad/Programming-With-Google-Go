package main

import "fmt"

func main() {
	var userStr float64
	fmt.Printf("Your string: ")
	fmt.Scan(&userStr)

	fmt.Printf("%.0f\n", userStr)
}
