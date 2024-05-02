package main

import (
	"fmt"
)

func main() {
	/*Em go a unica estrutura de repetição em go é o FOR*/

	a, b := 0, 10

	for a < b {
		a += 1
	}

	// for com a cláusula de inicialização

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println(a)

	// forma mais comun de interar sobre slices e utilizando o operador RANGE

	for indice, valor : slice {}
}
