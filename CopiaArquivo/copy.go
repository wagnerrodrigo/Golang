package main

import (
	"io"
	"log"
	"os"
)

func main() {
	origFile, err := os.Open("arquivo_original.txt")
	if err != nil {
		log.Fatal("erro ao abrir o arquivo")
	}
	defer origFile.Close()

	destFile, err := os.Create("arquivo_destino.txt")
	if err != nil {
		log.Fatal("erro ao ")
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, origFile)
	if err != nil {
		log.Fatal("erro ao copiar arquivo")
	}

}
