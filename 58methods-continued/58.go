package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func Aritmetica(number MyFloat) float64{
	if number < 0 {
		return float64(-number)
	}
	return float64(number)
}

func (number MyFloat) Abs() float64{
	if number < 0 {
		return float64(-number)
	}
	return float64(number)
}

func main(){
	numero := MyFloat(-math.Sqrt2)
	fmt.Println("1: ", numero) 
	fmt.Println("2: ", Aritmetica(numero))
	fmt.Println("3: ", numero.Abs())
}