// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/wagnerrodrigo/fortlev_sdk/fortlev/fortlev"
// )

// // https://api-platform.fortlevsolar.app/partner/user/sign-up
// func main() {
// 	client := fortlev.NewClient("https://api-platform.fortlevsolar.app", "")

// 	// Autenticação (caso necessário)
// 	err := client.Authenticate("energiasolardobrasil2018@gmail.com", "Esdb@2024")
// 	if err != nil {
// 		log.Fatalf("Erro ao autenticar: %v", err)
// 	}

// 	// Obter informações do usuário autenticado
// 	// userInfo, err := client.GetUserInfo()
// 	// if err != nil {
// 	// 	log.Fatalf("Erro ao obter informações do usuário: %v", err)
// 	// }

// 	fmt.Println("Informações do usuário:", client)
// }

package main

import (
	"fmt"
	"log"

	"github.com/wagnerrodrigo/fortlev_sdk/fortlev/fortlev"
)

func main() {
	// Inicializa o cliente Fortlev
	client := fortlev.NewClient("https://api-platform.fortlevsolar.app", "")

	// Autenticação
	err := client.Authenticate("energiasolardobrasil2018@gmail.com", "Esdb@2024")
	if err != nil {
		log.Fatalf("Erro ao autenticar: %v", err)
	}

	// userMe, err := client.GetUserInfo()
	// if err != nil {
	// 	log.Fatalf("Erro ao obter minhas informações: %v", err)
	// }

	// Mostrando o token obtido
	fmt.Println("Token de autenticação obtido:", client.AuthToken)

	// userId := 3
	// userInfo, err := client.GetUserId(userId)
	// if err != nil {
	// 	log.Fatalf("Error ao obter informações do usuário: %v", err)
	// }

	// fmt.Println("Token de autenticação obtido:", userId, userInfo)

	response, err := client.GetCities(2, 1)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Printf("Página atual: %d\n", response.CurrentPage)
	fmt.Printf("Cidades por página: %d\n", response.DocsPerPage)

	// for _, city := range response.Cities {
	// 	fmt.Printf("Cidade: %s - Estado: %s (%s)\n",
	// 		city.Name, city.StateName, city.StateInitials)
	// }

}
