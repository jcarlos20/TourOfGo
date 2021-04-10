package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("When is saturday? ")
	today := time.Now().Weekday()
	fmt.Println("Today: ", today)
	
	fmt.Println("Time Saturday: ", time.Saturday)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}