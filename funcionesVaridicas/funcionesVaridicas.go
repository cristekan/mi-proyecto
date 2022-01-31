package main

import "fmt"

//nombre de la funcion Variable Entrada int fuera del parentesis tipo de retorno
func sumar(numeros ...int) int {
	// el total inicial es 0
	total := 0
	// recorrer todos los numeros el guin bajo se utiliza como identificador de espacio en blanco
	for _, numero := range numeros {
		// en cada iteración sumar al total el valor del numero
		total = numero + total
	}
	// retornar el valor total
	return total
}

func main() {
	fmt.Println(sumar(2))
	fmt.Println(sumar(2, 2))
	fmt.Println(sumar(5, 4, 3))
}
