package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var urlStore = make(map[string]string)

func shorten(w http.ResponseWriter, r *http.Request) {
	// analisar os dados enviado no corpo a requisição
	r.ParseForm()
	// recuperar o valor do paramentro url do formulario
	url := r.Form.Get("url")
	// Gera um hass MD% da URL e usa os primeiros 5 caracteres como a URL encurtada.
	shortUrl := fmt.Sprintf("%x", md5.Sum([]byte(url)))[:5]
	// armazena no map urlStore a associação entre a URL encurtada e a URL original
	urlStore[shortUrl] = url
	// responde com o camninho da URL encurtada, usando o fmt.Fprintf
	fmt.Fprintf(w, "http://localhost:8080/%s\n", shortUrl)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	// Manipula requisições Get para o caminho /{shortUrl}
	// Usa mux.Vars para extrair o valor do parâmentro shortUrl da URL
	vars := mux.Vars(r)
	// Verifica se o shortUrl está no map urlStore
	originalUrl, ok := urlStore[vars["shortUrl"]]
	if ok {
		// Se Exixtir, redireciona para a Url original com http.Redirect
		http.Redirect(w, r, originalUrl, http.StatusMovedPermanently)
	} else {
		// Caso contrário, retorna um error 404 com http.NotFound
		http.NotFound(w, r)
	}
}

func main() {
	// Inicia http router
	r := mux.NewRouter()
	// roteamento
	r.HandleFunc("/create", shorten).Methods("POST")
	r.HandleFunc("/{shortUrl}", redirect).Methods("GET")
	// Inicia o servido na porta 8080
	log.Fatal(http.ListenAndServe(":8080", r))

}
