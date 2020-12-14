package main

import "fmt"

func main() {
	/*
		definir una variable y asignarle el tipo y se puede
		reasignar su valor
	*/
	var nombre string = "ian"
	fmt.Println(nombre)
	nombre = "Mauricio"
	fmt.Println(nombre)

	/*
		declaracion de multiples variables
	*/
	var a, b int = 10, 4
	fmt.Println(a, b)

	/*
		definir varibles de distinto tipo
	*/

	var (
		pi        float64 = 3.141592
		respuesta bool    = true
		cadena            = "esto des un string sin tipado"
		edad              = 80
	)
	fmt.Println(pi, respuesta, cadena, edad)

	/*
		otra forma de declarar una variable y asignarle valor
		al definir la variable se puede con :=
	*/

	vi := "esto es una variable tambien"
	fmt.Println(vi)
}
