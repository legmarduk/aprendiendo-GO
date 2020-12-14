package main

import "fmt"

const nombre string = "ian lopez"

func main() {

	/*
		se pueden declar dentro o fuera de una funcion main y
		su valor no puede ser modificado
	*/
	const n = 25

	fmt.Println(n, nombre)
}
