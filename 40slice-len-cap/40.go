package main

import "fmt"

func main(){
	s := []int{2,3,5,7,11,13}
	printSlices(s)

	// Slice the slice to gice it zero length
	s = s[:0]
	printSlices(s)

	// Extend its length
	s = s[:4]
	printSlices(s)

	// Drop its first two values
	s = s[2:]
	printSlices(s)
}

func printSlices(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}