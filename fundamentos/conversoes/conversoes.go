package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota) // Tira a casa decimal (arrenda para baixo)
	fmt.Println(notaFinal)

	// CUIDADO!
	fmt.Println("Teste " + string(97)) // convert to rune

	// int para string
	fmt.Println("Teste " + strconv.Itoa(97))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// string para boolean

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}

}
