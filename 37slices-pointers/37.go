package main

import "fmt"

func main(){
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("1", names)
	a := names[0: 2]
	b := names[1: 3]
	fmt.Println("2", a, b)

	b[0] = "XXX"
	fmt.Println("3", a, b)
	fmt.Println("4", names)
}