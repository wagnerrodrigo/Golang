package main

import (
	"log"

	"github.com/atotto/clipboard"
	"github.com/gen2brain/beeep"
	"github.com/google/uuid"
)

func main() {
	// cria o uuid
	newUUID := uuid.New()

	// converte o uuid para string
	uuidString := newUUID.String()

	// uuid com aspas
	// uuidWithQuotes := fmt.Sprintf("'%s'", uuidString)

	// colocar uuid na área de transferencia
	err := clipboard.WriteAll(uuidString)

	if err != nil {
		log.Fatalf("Erro ao copiar para a área de transferencia: %v", err)
	}

	// fmt.Printf("UUID: %s", uuidString)

	err = beeep.Notify("UUID:", uuidString, "uuid.jpg")
	if err != nil {
		log.Panicf("Erro ao notificar: %v", err)
	}

}
