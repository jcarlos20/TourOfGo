package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		fmt.Println("1- v:", v)
		fmt.Println("1- lim", lim)
		return v
	}
	fmt.Println("2- lim", lim)
	return lim
}

func main(){
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 2, 20),
	)
}