package main

import "fmt"

func main() {
	grades := []int32{4, 84, 29, 57}

	for i, grade := range grades {
		if grade > 37 {
			mod := grade % 10
			if mod >= 6 {
				if 10-mod < 3 {
					grade = grade - mod + 10
				}
			} else if mod <= 5 {
				if 5-mod < 3 {
					grade = grade - mod + 5
				}
			}
		}

		grades[i] = grade
	}

	fmt.Print(grades)
}
