package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}

func (object Vertex) Abs() float64 {
	return math.Sqrt(object.X*object.X + object.Y*object.Y)
}

func (object *Vertex) Scale(number float64){
	object.X = object.X * number
	object.Y = object.Y * number
}

func main(){
	objeto := Vertex{3,4}
	fmt.Println("Objeto: ", objeto)
	objeto.Scale(10) 
	fmt.Println("Objeto Scale: ", objeto)
	fmt.Println("Abs: ", objeto.Abs()) 
}