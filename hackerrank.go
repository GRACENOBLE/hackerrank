/* this program creates a function that takes an array and a constant and finds the median of the array then tries to make the median equal the constant*/ 
package main

import (
	"fmt"
	"sort"
)

func median(numbers []float64) float64 {
	sort.Float64s(numbers)
	length := len(numbers)

	if length%2 == 0 {
		return (numbers[length/2-1] + numbers[length/2]) / 2
	}

	return numbers[length/2]
}

func operations(arr []float64, con float64) float64 {

	med := median(arr)
	var count float64

	for med != con {
		if med < con {
			med++
			count++
		} else if med > con {
			med--
			count++
		}
	}
	return count

}

func main() {
	arr := []float64{1, 2, 3, 4, 5, 6}

	fmt.Print(median(arr))

	// fmt.Printf("The number of operations is %v", operations(arr, 5))

}
