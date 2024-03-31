package main

import (
	"errors"
	"fmt"
)

func main() {
	pilha := Pilha{}
	fmt.Println("Pilha criada com tamanho ", pilha.Tamanho())

	pilha.Empilhar("GO")
	pilha.Empilhar(2020)
	pilha.Empilhar("Wagner")
	pilha.Empilhar(3.14)
	fmt.Println("Tamanho após empilhar 4 valores: ", pilha.Tamanho())
	fmt.Println("Vazia? ", pilha.Vazia())

	for !pilha.Vazia() {
		v, _ := pilha.Desempilhar()
		fmt.Println("Desempilando ", v)
		fmt.Println("Tamanho ", pilha.Tamanho())
		fmt.Println("Vazia? ", pilha.Vazia())
	}

	_, err := pilha.Desempilhar()
	if err != nil {
		fmt.Println("Erro: ", err)
	}
}

type Pilha struct {
	valores []interface{}
}

func (pilha Pilha) Tamanho() int {
	return len(pilha.valores)
}

func (pilha Pilha) Vazia() bool {
	return pilha.Tamanho() == 0
}

func (pilha *Pilha) Empilhar(valor interface{}) {
	pilha.valores = append(pilha.valores, valor)
}

func (pilha *Pilha) Desempilhar() (interface{}, error) {
	if pilha.Vazia() {
		return nil, errors.New("Pilha vazia!")
	}
	valor := pilha.valores[pilha.Tamanho()-1]
	pilha.valores = pilha.valores[:pilha.Tamanho()-1]
	return valor, nil
}
