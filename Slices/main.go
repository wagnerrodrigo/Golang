package main

import (
	"fmt"
)

func main() {
	var a []int
	primos := []int{2, 3, 5, 7, 11, 13}
	nomes := []string{}
	fmt.Println(a, primos, nomes)

	/*
		na prática slices é muito similar ao array
	*/

	/*um slice pode ser criado atavés da função make(), que internament aloca um array e retorna uma referência para o slice
	criado.
	Sua assinatura;

	fun make([]T, len, cap) []T

	T -> representa o tipo dos elementos do slice,
	len -> o tamanho inicial do array alocado
	cap -> o tamanha total da área de memória reservada para o crescimento do silce, esse argumento e omitido por
	conveniência e o Go assume o padrão o mesmo valor do tamanho.
	*/
	// forma de criar um Slice
	b := make([]int, 10)
	fmt.Println(b, len(b), cap(b))

	c := make([]int, 10, 20)
	fmt.Println(c, len(c), cap(c))

	/*A vantagem de utilizar slices de fato, qualquer tipo criado através da função make()- em vez de arrays é que, quando usados como
	argumentos ou no retorno de funções, são passados por referência e não por cópia, iso torna estas chamadas muito mais eficientes, pois o tamanho
	da referência será sempre o mesmo, independente do tamanho do slice.*/

}
