package main

import "fmt"

// por padrão, BREAK sai do loop mais proximo ao ponto que o comando foi executado.
// Há casos nos quais temos LOOPS For aninhados e desejamos quebrar o loop externo
// em vez do interno. Para resolver este problema, Go também dá suporte a blocos
// For nomeados.

var i int

func main() {

externo:
	for {
		for i = 0; i < 10; i++ {
			if i == 5 {
				break externo
			}
		}
	}
	fmt.Println(i)

}
