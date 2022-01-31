package main

import "fmt"

const pi = 3.1416

func circulo(radio float64) (area float64, perimetro float64) {
	area = pi * radio * radio
	perimetro = 2 * pi * radio
	return
}

func main() {

	var radio float64

	fmt.Printf("Ingrese un nuevo valor para el radio")
	fmt.Scanf("%f", &radio)

	area, perimetro := circulo(radio)

	fmt.Println("El area es:", area)
	fmt.Println("el perimetro es:", perimetro)
}
