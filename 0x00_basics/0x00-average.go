package main

import "fmt"

func average(numbers ...float32) float32 {
	var result float32 = 0.0
	for _, number := range numbers {
		result += number
	}
	result = result / float32(len(numbers))
	return result
}

func main() {
	fmt.Printf("Average %f", average(5.2, 5.3, 3.2, 3.7, 9.3))
}
