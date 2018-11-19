package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	result := 0.04*math.Pow(x, 3) + math.Pow(x, 2) - 7*x - 7
	return result
}

func esmMin(x0 float64, h float64, kMax int) {
	fmt.Println("\n\t\tEven Search Method to find the minimum")
	var x1 float64
	k := 0
	yf0 := f(x0)
	yf1 := f(x0 + h)
	for k < kMax {
		k++
		x1 = x0 + h
		if yf1 >= yf0 {
			x1 = x0
			yf1 = yf0
		} else {
			x0 = x1
			yf0 = yf1
			x1 = x0 + h
			yf1 = f(x1)
		}
	}

	fmt.Println("x1:", x1)
	fmt.Println("yf1:", yf1)
	fmt.Println("k:", k)
	fmt.Println("h:", h)
}

func esmMax(x0 float64, h float64, kMax int) {
	fmt.Println("\n\n\t\tEven Search Method to find the maximum")
	var x1 float64
	k := 0
	yf0 := f(x0)
	yf1 := f(x0 + h)
	for k < kMax {
		k++
		x1 = x0 + h
		if yf1 <= yf0 {
			x1 = x0
			yf1 = yf0
		} else {
			x0 = x1
			yf0 = yf1
			x1 = x0 + h
			yf1 = f(x1)
		}
	}

	fmt.Println("x1:", x1)
	fmt.Println("yf1:", yf1)
	fmt.Println("k:", k)
	fmt.Println("h:", h)
}

func main() {
	var x0 float64
	var h float64
	kMax := 6
	fmt.Print("Enter x0: ")
	fmt.Scanln(&x0)
	fmt.Print("Enter h: ")
	fmt.Scanln(&h)

	esmMin(x0, h, kMax)
	esmMax(x0, h, kMax)
}
