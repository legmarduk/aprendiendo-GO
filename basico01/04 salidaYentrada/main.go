package main

import "fmt"

/*
ej:
	nombre :="ian"
	edad :=27

	fmt.Println("Nombre", nombre, "Edad", edad)
	fmt.Printlf("Nombre %s Edad %d",nombre, edad) como en C

*/

func main() {

	var (
		nombre string
		edad   int
	)

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanln(&nombre) // otra forma fmt.Scanf("%s", &nombre)
	fmt.Print("Ingresa tu edad: ")
	fmt.Scanln(&edad)

	fmt.Println("Tu nombre es", nombre, "y tu Edad es", edad)
	//fmt.Printf("Tu nombre es %s y tu Edad es  %d", nombre, edad)

}
