package main

import (
	"fmt"

	rotarcadenas "github.com/Zumari/labora-api/ejercicios/ADN/rotarCadenas"
)

func main() {
	s := "ABCD"
	a := "DABC"

	fmt.Println(len(s))
	fmt.Println("Cadena original: ", rotarcadenas.PrintCadena(s))
	fmt.Println("Cadena rotada a la derecha: ", rotarcadenas.RotarDerecha(s))
	fmt.Println("Cadena rotada a la izquierda: ", rotarcadenas.RotarIzquierda(s))

	rotarcadenas.VerificarADN(s, a)

}
