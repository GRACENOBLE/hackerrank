package main

import "fmt"

func main() {

	arr := []int{1, 1, 0, -1, -1}
	negatives := 0
	zeros := 0
	positives := 0

	for _, item := range arr {

		if item < 0 {
			negatives++
		} else if item == 0 {
			zeros++
		} else {
			positives++
		}
	}

	negativeRatio := float64(negatives) / float64(len(arr))
	positiveRatio := float64(positives) / float64(len(arr))
	zeroRatio := float64(zeros) / float64(len(arr))

	fmt.Printf("%.6f\n%.6f\n%.6f", negativeRatio, positiveRatio, zeroRatio)
}
