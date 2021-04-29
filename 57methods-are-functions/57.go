package main

import (
	"fmt"
	"math"
)

type Coordinates struct{
	X, Y float64
}

func Abs(object Coordinates) float64{
	return math.Sqrt(object.X * object.X + object.Y*object.Y)
}

func main(){
	object := Coordinates{3,4}
	fmt.Println(object);
	fmt.Println(Abs(object))
}