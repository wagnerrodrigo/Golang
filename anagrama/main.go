package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isAnagram("banana", "anbana"))
}

func isAnagram(s string, t string) bool {
	// criar um mapa
	anagramMap := make(map[string]int) // inicializando

	// var anagramMap map[string]int dessa forma daria erro no for pois não estou inicializando o map
	// var anagramMap map[string]int{} inicializando o map vazio

	// split para transforma a palavra em um slice
	s1 := strings.Split(s, "")

	// adicinar cada letra com chave do map e soma
	for _, char := range s1 {
		anagramMap[char]++

		// verifica se a chave já existe, se sim soma
		// if _, ok := anagramMap[char]; ok {
		// 	anagramMap[char]++
		// } else {
		// 	// se não existir cria a chave no map
		// 	anagramMap[char] = 1
		// }
		// ou simplimente anagramMap[char]++ removendo tudo pois o map são iniciado com zero
	}

	t1 := strings.Split(t, "")
	for _, char := range t1 {
		// verifica se a chave já existe, se sim subitrair
		anagramMap[char]--
	}

	fmt.Println(anagramMap)
	for _, value := range anagramMap {
		if value != 0 {
			return false
		}
	}

	return true
}
