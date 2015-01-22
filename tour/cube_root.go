package main

import (
	"fmt"
	. "math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := 4 + 0i
	for i := 0; i < 10; i++ {
		z = z - (Pow(z, 3)-x)/(3*Pow(z, 2))
	}
	return z

}

func main() {
	fmt.Println(Cbrt(27))

}
