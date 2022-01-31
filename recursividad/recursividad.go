package main

import "fmt"

func correr(corredor int, totaldecorredor int) {
	fmt.Println("Corredor", corredor, "arranca la carrera")
	bandera := totaldecorredor > corredor

	if bandera {
		siguienteCorredor := corredor + 1
		fmt.Println("corredor", corredor, "le dio la partida al siguiente", siguienteCorredor)
		correr(siguienteCorredor, totaldecorredor)
	}
	fmt.Println("el corredor", corredor, "termino de correr")
}

func main() {
	correr(1, 3)
}
