package main

import "fmt"

func circulo(radio float64) (area func() float64, perimetro func() float64) {

	area = func() float64 {
		return 3.1416 * radio * radio
	}

	perimetro = func() float64 {
		return 2 * 3.1416 * radio
	}

	return
}

func main() {
	area, perimetro := circulo(10)
	fmt.Println("El area del circulo es", area())
	fmt.Println("El perimetro del circulo es", perimetro())

}
