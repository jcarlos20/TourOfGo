package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (object Vertex) Abs() float64 {
	return math.Sqrt(object.X * object.X + object.Y * object.Y)
}

func main() {
	object := Vertex{3,4}
	fmt.Println(object.Abs())
}