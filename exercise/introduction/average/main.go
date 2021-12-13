package main

import "fmt"

func main() {
	nums := []float64{10, 50, 2.3, 0.2, 43.2, 2.2}

	sum := 0.0
	var average float64

	for _, n := range nums {
		sum += n
	}

	average = sum / float64(len(nums))
	fmt.Println(average)

}
