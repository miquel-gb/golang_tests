package main

import (
	"fmt"
	"math"
)

func Pic(dx, dy int) [][]uint8 {
	// initialize the image
	image := make([][]uint8, dy)

	// iterate over each pixel
	for x := 0; x < dy; x++ {
		image[x] = make([]uint8, dx)
		for y := 0; y < dx; y++ {
			// set a color value for each pixel
			image[x][y] = uint8((x ^ y) * (x ^ y))

		}
	}
	return image
}

func Sqrt(x float64) float64 {
	z := 1.0
	prevZ := z
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		dif := math.Abs(prevZ - z)
		/* Control print */
		fmt.Printf("X = %v, Z = %v, Z^2 = %v, prevZ = %v, dif = %v\n", x, z, z*z, prevZ, dif)

		if dif < 1e-15 {
			return z
		}
		prevZ = z
	}
	return z
}

func forAsWhile() {
	i := 1
	for i < 10 {
		fmt.Printf("Vuelto número %v\n", i)
		i++
	}
}

func pow(x, n, lim float64) float64 {
	/* Functin before if condition, only visible for if scope */
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	/*fmt.Println(Pic())*/
	/*forAsWhile()*/

}
