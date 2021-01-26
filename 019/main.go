package main

import (
	"fmt"
	"time"
)

// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

func main() {

	start := time.Date(1901, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC)

	var daysAdd int
	var sundays int
	for {
		daysAdd++
		testDate := start.AddDate(0, 0, daysAdd)
		//		fmt.Printf("%v\n", testDate)

		if int(testDate.Weekday()) == 0 && int(testDate.Day()) == 1 {
			sundays++
		}

		if testDate.After(end) {
			break
		}
	}
	fmt.Println("019/ Sundays:", sundays)
}
