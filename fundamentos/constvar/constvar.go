package main

import (
	"fmt"
	m "math" // É possivel usar um label para os pacotes
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo float64 -> inferido pelo compilador

	// forma reduzida de se criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("Área da circunferencia:", area)

	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false // Setar duas variaveis ao mesmo tempo
	fmt.Println(e, f)

	g, h, i := 2, false, "opa" // Go é fortemente tipado
	fmt.Println(g, h, i)
}
