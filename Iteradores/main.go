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

	fmt.Println("rst de a ", a)

	// forma mais comun de interar sobre slices e utilizando o operador RANGE
	// for indice, valor : slice {}

	numeros := []int{1, 2, 3, 4, 5}

	for i := range numeros {
		numeros[i] *= 2
	}

	fmt.Println("\n", numeros)

	// o código anterior itera sobre um slice chamado numeros, multiplicando cada valor por 2. após a interação imprimemo o conteúdo de numeros e obtemos [2 4 6 8 10]

	// Quando não precisamos dos indices dos valores, podemos ignorá-los ao identificador vazio _
	// for _, elemento := range silece {}

}
