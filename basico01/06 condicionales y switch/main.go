package main

import "fmt"

func main() {

	var n int
	fmt.Println("Ingresa un numero entre 1 y 5")
	fmt.Print("Ingrese un numero: ")
	fmt.Scanln(&n)

	if n == 0 {
		fmt.Println("Ingresaste un 0")
		return
	} else if n > 0 && n < 6 {
		fmt.Println("Ingresaste un numero dentro del rango")
	} else {
		fmt.Println("Ingresaste un numero fuera del rango")
		return
	}

	switch n {
	case 1:
		fmt.Println(" Ingresaste un uno")
		break
	case 2:
		fmt.Println(" Ingresaste un dos")
		break
	case 3:
		fmt.Println(" Ingresaste un tres")
		break
	case 4:
		fmt.Println(" Ingresaste un cuatro")
		break
	case 5:
		fmt.Println(" Ingresaste un cinco")
		break
	default:
		fmt.Println(" Ingresaste otra cosa")
	}

}
