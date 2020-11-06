/*package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return a*math.Pow(t, 2)/2 + v0*t + s0
	}
}

func main() {
	var a, v0, s0, t float64

	fmt.Print("Enter value for acceleration - ")
	_, err := fmt.Scanf("%f", &a)
	if err != nil {
		fmt.Print("Incorrect input\n")
		return
	}
	fmt.Print("Enter value for initial velocity - ")
	_, err = fmt.Scanf("%f", &v0)
	if err != nil {
		fmt.Print("Incorrect input\n")
		return
	}
	fmt.Print("Enter value for initial displacement - ")
	_, err = fmt.Scanf("%f", &s0)
	if err != nil {
		fmt.Print("Incorrect input\n")
		return
	}
	DispByTime := GenDisplaceFn(a, v0, s0)
	fmt.Print("Enter value for time - ")
	_, err = fmt.Scanf("%f", &t)
	if err != nil {
		fmt.Print("Incorrect input\n")
		return
	}
	fmt.Printf("\nDisplacement after %f seconds: ", t)
	fmt.Print(DispByTime(t))
	fmt.Print("\n")
}
*/