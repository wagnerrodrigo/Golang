package main

import (
	"fmt"
	m "math" // a letra M e uma forma de renomear a importação
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)

	fmt.Println("A área da circunfeêrencia é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)
	var e, f bool = true, false // o true vai para a variavel e e o false vai para a variavel f
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!" // O GO E UMA LIGUAGUEM STATICAMENTE TIPADA E FORTMENTE TIPADA, TODOAS AS VARIAVEIS TEM UM TIPO E ESSE TIPO NÃO VARIA DURATE A EXECUÇÃO
	fmt.Println(g, h, i)
}
