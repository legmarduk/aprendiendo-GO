package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	/*
		tecnicamente no tenemos un ciclo while
		pero podemos usar un for para eso
	*/
	fmt.Println("nuestro 'while'")
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	/*
		tampoco tenemos una ciclo do while tenemos la opcion
		de usar el for
	*/
	fmt.Println("nuestro 'do while'")
	do := 0
	for {
		fmt.Println(do)
		if !(do < 10) {
			break
		}
		do++
	}
	/*
		no tenemos un foreach pero tambien puede ser remplazado
		por un for
	*/
	fmt.Println("nuestro 'foreach'")
	lista := []string{"otto", "ian", "huachipato", "tumama"}

	//for{key},{value} := range{list}
	//foreach sin key
	for _, dato := range lista {
		fmt.Println("=>", dato)
	}

	//foreach con key

	animales := map[string]string{
		"perro": "woof",
		"gato":  "meow",
		"pato":  "cuac",
	}

	for animal, ruido := range animales {
		fmt.Println("el", animal, "hace", ruido)
	}
}
