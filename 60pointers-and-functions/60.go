package main

import (
	"fmt"
	"math"
)

type Coordenadas struct {
	X, Y float64
}

func Abs(objeto Coordenadas) float64{
	return math.Sqrt(objeto.X*objeto.X + objeto.Y*objeto.Y)
}

func Scale(objeto *Coordenadas, numero float64){
	objeto.X = objeto.X * numero
	objeto.Y = objeto.Y * numero
}

func main(){
	objeto := Coordenadas{3, 4}
	fmt.Println("Objeto: ", objeto)
	Scale(&objeto, 10)
	fmt.Println("Objeto Scale: ", objeto)
	fmt.Println("Objeto Abs: ", Abs(objeto))
}