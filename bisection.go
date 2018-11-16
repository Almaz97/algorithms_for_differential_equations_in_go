// 0.04 * x ** 3 + x ** 2 - 7 * x - 7, 4

package main

import (
	"fmt"
	"math"
	"os"
)

func f(x float64) float64 {
	result := 0.04*math.Pow(x, 3) + math.Pow(x, 2) - 7*x - 7
	return result
}

func bisection(a float64, b float64, tol float64, kMax int) {
	fmt.Println("\n\n\t\tBisection Search Method")
	fa := f(a)
	fb := f(b)

	if math.Signbit(fa) == math.Signbit(fb) {
		fmt.Println("Sign of f(a) and f(b) must be opposite! Check endpoints of the internal [a,b]")
		os.Exit(1)
	}

	k := 0

	for b-a > tol && k < kMax {
		m := a + (b-a)/2
		fa = f(a)
		fb = f(m)
		if math.Signbit(fa) == math.Signbit(fb) {
			a = m
		} else {
			b = m
		}
		k++
	}

	fmt.Println("a: ", a)
	fmt.Println("b:", b)
	fmt.Println("F(a):", fa)
	fmt.Println("F(b):", fb)
	fmt.Println("K: ", k)
}
func main() {
	var a float64
	var b float64
	var tol float64
	var kMax int
	fmt.Print("Enter a: ")
	fmt.Scanln(&a)
	fmt.Print("Enter b: ")
	fmt.Scanln(&b)
	fmt.Print("Enter tol: ")
	fmt.Scanln(&tol)
	fmt.Print("Enter kMax: ")
	fmt.Scanln(&kMax)

	bisection(a, b, tol, kMax)

}
