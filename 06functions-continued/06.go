package main

import "fmt"

func add(x,y,z int) int {
	return x + y + z
}

func main(){
	fmt.Println(add(5, 10, 15))
}