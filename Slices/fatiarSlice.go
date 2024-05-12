package main

import "fmt"

func main() {

	//Fatiar(do ingles to slice) é o nome dado a operações que extraem partes de um slice ou de um array.
	// para fazer um slice ou array, utilizamos a seguinte forma:
	// novoSlice := slice[inicio : fim]
	// Sendo que 0 <= inicio <= fim <= len(slice)
	// qualquer combinação de índices inicial e final que não atenda a esta regra resulta em um erro
	// de compilação (SLICE BOUNDS OUT OF RANGE, OU LIMITES DO SLICE FORA DE ALCANCE)

	// fib := []int{1, 1, 2, 3, 5, 8, 13}
	// fmt.Println(fib)
	// fmt.Println(fib[:3])
	// fmt.Println(fib[2:])
	// fmt.Println(fib[:])

	original := []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", original)

	novo := original[1:3]
	fmt.Println("Novo:", novo)

	original[2] = 13

	fmt.Println("Original pós modificação:", original)
	fmt.Println("Novo pós modificação:", novo)

}
