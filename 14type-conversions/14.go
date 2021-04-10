package main

import (
	"fmt"
	"math"
)

func main(){
	var x, y int = 3, 4
	fmt.Println(float64(x*x))
	fmt.Println(float64(y*y))
	fmt.Println(float64(x*x + y*y))
	
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println(f)
	var z uint = uint(f)
	fmt.Println(z)
	fmt.Println(x, y, z)
}