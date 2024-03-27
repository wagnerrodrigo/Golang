package main

import "fmt"

func imprimirDados(nome string, idade int) { // função sem retorno
	fmt.Printf("%s tem %d anos. \n", nome, idade)
}

func soma(n, m int) int { // Forma simplificada de atribuir o mesmo tipo para os dois argumentos (N e M são int)
	return n + m // função com retorno
}

func main() {
	imprimirDados("Arthur Dent ", 34)
	s := soma(21, 21)
	fmt.Println("A soma é", s)
}
