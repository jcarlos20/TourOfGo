package main

import "fmt"

type Coordenadas struct {
	X, Y float64
}

func (objeto *Coordenadas) Scale(number float64){
	objeto.X = objeto.X * number
	objeto.Y = objeto.Y * number
}

func ScaleFunc(objeto *Coordenadas, number float64){
	objeto.X = objeto.X * number
	objeto.Y = objeto.Y * number
} 

func main(){
	objeto := Coordenadas{3, 4}
	fmt.Println("Objeto: ", objeto)
	objeto.Scale(2)
	fmt.Println("Objeto Scale: ", objeto)
	ScaleFunc(&objeto, 10)
	fmt.Println("Objeto ScaleFunc: ", objeto)
	
	objeto2 := &Coordenadas{3, 4}
	fmt.Println("Objeto2: ", objeto2)
	objeto2.Scale(3)
	fmt.Println("Objeto2 Scale: ", objeto2)
	ScaleFunc(objeto2, 8)
	fmt.Println("Objeto2 ScaleFunc: ", objeto2)

	fmt.Println("Objeto1: ", objeto, " ", "Objeto2: ", objeto2);
}