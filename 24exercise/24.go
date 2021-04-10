package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	//fmt.Printf("Sqrt approximation of %v:\n", x)
	//for i := 1; i <=10; i ++ {
	//	z -= (z*z - x) / (2*z)
	//	fmt.Printf("Iteration %v, value = %v\n", i, z)		
	//}
	var t float64
	for {
		//z, t = z - (z*z -x) / (2*z), z
		t = z
		z = z - (z*z -x) / (2*z)
		// fmt.Printf("Z: %v, T: %v\n", z, t);
		// fmt.Printf("DIF: %v\n", math.Abs(t-z))
		if math.Abs(t-z) < 1e-6 {
			break
		}
	}

	return z
}

func main(){
	guess := Sqrt(2)
	expected := math.Sqrt(2) 
	//fmt.Println(Sqrt(2))
	fmt.Printf("Guess: %v, Expected: %v, Error: %v", guess, expected, math.Abs(guess - expected))
}