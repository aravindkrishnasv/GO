package main

import "fmt"

func main() {
	var a float64 = 9
	var b float64 = 3
	var q = division(a,b)

	fmt.Printf("Value is : %.2f\n",q)
}

func division(num1,num2 float64) float64 {
	return num1 / num2
}