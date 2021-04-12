package main

import "fmt"

func main(){
	pow := make([]int, 10)
	for i := range pow{
		fmt.Printf("%v\n", i)
		pow[i] = 1 << uint(i)
	}
	fmt.Println("POW: \n", pow)

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}