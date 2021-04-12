package main

import "fmt"

var test = []int{1,2,4,8,16,32,64,128}

func main(){
	for i, v := range test {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for x := 0; x < len(test); x++ {
		fmt.Println(test[x])
	}
}