package main

import (
	"fmt"
)

func main() {
	var a [3]int // zero value em go
	numeros := [5]int{1, 2, 3, 4, 5}
	primos := [...]int{2, 3, 5, 7, 11, 13}
	nomes := [2]string{}

	/*Este valor inicial varia de acord com o tipo de dados definido para o array, da seguite forma:
	- false para valores do tipo bool;
	- 0 para ints;
	- 0.0 para floats;
	- ""(ou string vazia) para strings;
	- nil para ponteiros, funções, interfaces, slices, maps e channels.

	no caso do primos, o nome do operador [...] ELLIPSIS(reticências) esse operador informa para o compilador a calcular o tamanho do array com base
	na quantidade de elementos declarados.

	*/

	fmt.Println(a, numeros, primos, nomes)
	fmt.Println("\nTamanho do array->>", len(a), len(numeros), len(primos), len(nomes))

	/*Array multidimensionais*/

	var multiA [2][2]int

	multiA[0][0], multiA[0][1] = 3, 5
	multiA[1][0], multiA[1][1] = 7, 2

	multiB := [2][2]int{{2, 13}, {-1, 6}}

	fmt.Println("\nmultiA:", multiA)
	fmt.Println("multiB:", multiB)

}
