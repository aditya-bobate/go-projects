package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var i int = 32
	var f float32 = 45.21
	var f2 float64 = 67.9598348795
	var s string = "Aditya"
	var b bool = true
	var c complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Println(i, f, f2, s, b, c)
}
