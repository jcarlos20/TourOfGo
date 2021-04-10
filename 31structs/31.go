package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type City struct {
	name string
	country string
}

func main(){
	// fmt.Println(Vertex{1, 2})
	fmt.Println(City{"Buenos Aires", "Argentina"})
}