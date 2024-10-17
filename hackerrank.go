package main

import (
	"fmt"
	"strconv"
)

func main(){
	s := "12:45:54PM"
	time := s[:len(s)-2]
	symbol := s[len(s)-2:]
	
	if symbol == "AM"{
		hour := time[:2]
		
		if hour == "12" {
			hour = "00"
		}
		
		print(hour+time[2:])
		return
	}

	hour, err := strconv.Atoi(s[:2])

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		
		if hour != 12{
			hour += 12
		}
	}

	print(strconv.Itoa(hour) + s[2:len(s)-2])

}