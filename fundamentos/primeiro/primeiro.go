// Programas executaveis iniciam pelo pacote main
package main

/*
Os códigos em Go são organizados em pacotes.
para usa-los é necessários declarar um ou mais imports
*/
import "fmt"

// A porta de entrada de um programa Go é a função main
func main() {
	fmt.Print("Hello ")
	fmt.Print("World!")
}
