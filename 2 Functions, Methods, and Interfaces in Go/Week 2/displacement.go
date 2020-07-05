package main

import "fmt"

func GetDisplaceFn(a, b, c float64) func(float64) float64 {
	fn := func(d float64) float64 {
		return (0.5)*(a*d*d) + (b * d) + c
	}
	return fn
}

func main() {
	var a, b, c, d float64
	fmt.Printf("Enter value for acceleration: ")
	fmt.Scanf("%f", &a)
	fmt.Printf("Enter value for intial velocity: ")
	fmt.Scanf("%f", &b)
	fmt.Printf("Enter value for intial displacement: ")
	fmt.Scanf("%f", &c)
	fmt.Printf("Enter the value of time: ")
	fmt.Scanf("%f", &d)

	fn := GetDisplaceFn(a, b, c)

	fmt.Printf("The displacement is: ")
	fmt.Println(fn(d))
}
