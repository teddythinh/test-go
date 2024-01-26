package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	Tobe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	Packages()

	Imports()

	ExportedName()

	fmt.Println(Add(1, 2))
	fmt.Println(Add2(1, 2))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var i, j int = 1, 2
	k := 3 // short assignment statement can be used in place of a var declaration with implicit type
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, k, c, python, java)

	fmt.Printf("Type: %T Value: %v\n", Tobe, Tobe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	var d int
	e := d
	fmt.Println(d, e)

	const Pi = 3.14
	fmt.Println(Pi)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
