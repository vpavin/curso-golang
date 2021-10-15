package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é do tipo", reflect.TypeOf(32000))

	// sem sinal (positivos) -> uint8 uint16 uint32 uint64
	var b byte = 255 // byte => uint8
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal (positivo e negativo) -> int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo de int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa o mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é ", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais -> float32 float64
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O literal de 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string == delimitadas com aspas duplas
	// não é possivel utilizar aspas simples
	s1 := "Olá meu nome é Vini"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// strings com multiplas linhas
	s2 := `Olá
	meu nome
	é Vini`
	fmt.Println(s2 + "!")
	fmt.Println("O tamanho da string é", len(s2))

	// nao existe char em Go
	char := 'a'
	fmt.Println("O tipo char é", reflect.TypeOf(char)) //int32
	fmt.Println(char)

}
