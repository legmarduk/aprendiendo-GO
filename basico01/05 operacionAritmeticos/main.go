package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("ingresa un numero: ")
	fmt.Scanln(&a)
	fmt.Print("ingresa otro numero: ")
	fmt.Scanln(&b)

	suma := a + b
	resta := a - b
	multi := a * b
	div := a / b
	modulo := a % b

	fmt.Println("\nla suma es:", suma)
	fmt.Println("la restar es:", resta)
	fmt.Println("la multiplicacion es: ", multi)
	fmt.Println("la division es: ", div)
	fmt.Println("el modulo es: ", modulo)

	/* operadores relacionales*/
	fmt.Println("el valor de a es igual a b?", a == b)
	fmt.Println("el valor de a es distinto a b?", a != b)
	fmt.Println("el valor de a es mayor a b?", a > b)
	fmt.Println("el valor de a es menor a b?", a < b)
}
