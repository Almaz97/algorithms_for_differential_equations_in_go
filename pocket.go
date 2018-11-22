package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	result := 0.04*math.Pow(x, 3) + math.Pow(x, 2) - 7*x - 7
	return result
}

func pocketMin(h0 float64, h1 float64, x0 float64, r float64, tol float64, kMax int) {
	fmt.Println("\n\t\tPocket Search Method to find the minimum")
	x1 := x0 + h1
	yf0 := f(x0)
	yf1 := f(x1)
	k := 0

	for k < kMax {
		k++
		if yf1 >= yf0 {
			fmt.Println(true)
			if math.Abs(h0) < tol/r {
				h1 = h0
				x1 = x0
				yf1 = yf0
			} else {
				h1 = (-h0) / r
				h0 = h1
			}
		} else {
			fmt.Println(false)
			h1 = h0
			x0 = x1
			yf0 = yf1
			x1 = x0 + h1
			yf1 = f(x1)
		}
	}

	fmt.Println("x1: ", x1)
	fmt.Println("yf1: ", yf1)
	fmt.Println("k: ", k)
	fmt.Println("h1: ", h1)
}

func pocketMax(h0 float64, h1 float64, x0 float64, r float64, tol float64, kMax int) {
	fmt.Println("\n\n\t\tPocket Search Method to find the maximum")
	x1 := x0 + h1
	yf0 := f(x0)
	yf1 := f(x1)
	k := 0

	for k < kMax {
		k++
		if yf1 <= yf0 {
			fmt.Println(true)
			if math.Abs(h0) < tol/r {
				h1 = h0
				x1 = x0
				yf1 = yf0
			} else {
				h1 = (-h0) / r
				h0 = h1
			}
		} else {
			fmt.Println(false)
			h1 = h0
			x0 = x1
			yf0 = yf1
			x1 = x0 + h1
			yf1 = f(x1)
		}
	}

	fmt.Println("x1: ", x1)
	fmt.Println("yf1: ", yf1)
	fmt.Println("k: ", k)
	fmt.Println("h1: ", h1)
}

func main() {
	var h0 float64
	var x0 float64
	var r float64
	var tol float64
	kMax := 6
	fmt.Print("Enter h0: ")
	fmt.Scanln(&h0)
	fmt.Print("Enter x0: ")
	fmt.Scanln(&x0)
	fmt.Print("Enter r: ")
	fmt.Scanln(&r)
	fmt.Print("Enter tol: ")
	fmt.Scanln(&tol)
	pocketMin(h0, h0, x0, r, tol, kMax)
	pocketMax(h0, h0, x0, r, tol, kMax)

}
