package fortlev

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// FortlevClient representa o cliente para interagir com a API da Fortlev Solar
type FortlevClient struct {
	BaseURL    string
	AuthToken  string
	HTTPClient *http.Client
}

// NewClient cria um novo cliente Fortlev
func NewClient(baseURL, authToken string) *FortlevClient {
	return &FortlevClient{
		BaseURL:   baseURL,
		AuthToken: authToken,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// CityResponse representa a estrutura da resposta da API de cidades
type CityResponse struct {
	CurrentPage int    `json:"current_page"`
	DocsPerPage int    `json:"docs_per_page"`
	Data        []City `json:"data"`
}

// City representa a estrutura de uma cidade
type City struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	StateID       int    `json:"state_id"`
	StateName     string `json:"state_name"`
	StateInitials string `json:"state_initials"`
}

// // authenticate realiza a autenticação na API e atualiza o token
// func (c *FortlevClient) Authenticate(username, password string) error {
// 	url := fmt.Sprintf("%s/partner/user/login", c.BaseURL)

// 	payload := map[string]string{
// 		"username": username,
// 		"password": password,
// 	}
// 	jsonPayload, _ := json.Marshal(payload)

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
// 	if err != nil {
// 		return err
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := c.HTTPClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return errors.New("falha na autenticação")
// 	}

// 	body, _ := ioutil.ReadAll(resp.Body)
// 	var result map[string]string
// 	json.Unmarshal(body, &result)

// 	token, exists := result["token"]
// 	if !exists {
// 		return errors.New("token não encontrado na resposta")
// 	}

// 	c.AuthToken = token
// 	return nil
// }

func (c *FortlevClient) Authenticate(username, password string) error {
	url := fmt.Sprintf("%s/partner/user/login", c.BaseURL)

	// Criando o payload no formato x-www-form-urlencoded
	payload := fmt.Sprintf(
		"grant_type=&username=%s&password=%s&scope=&client_id=&client_secret=",
		username, password,
	)

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(payload))
	if err != nil {
		return fmt.Errorf("erro ao criar requisição: %w", err)
	}

	// adiciona os headers necessários
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Executando a requisição
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()

	// Lidando com erros de status HTTP
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("falha na autenticação: status %d, resposta: %s", resp.StatusCode, string(body))
	}

	// Processando a resposta para obter o access_token
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("erro ao ler a resposta: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("erro ao parsear resposta JSON: %w", err)
	}

	token, exists := result["access_token"].(string)
	if !exists {
		return errors.New("access_token não encontrado na resposta")
	}

	// Atualizando o token de autenticação do cliente
	c.AuthToken = token
	return nil
}

func (c *FortlevClient) GetUserId(userId int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/partner/user/%d", c.BaseURL, userId)

	// Cria a requisição HTTP GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json") // Opcional
	req.Header.Set("Authorization", "Bearer "+c.AuthToken)

	// Executa a requisição
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()

	// Log do status e corpo da resposta
	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("Status: %d, Resposta: %s", resp.StatusCode, string(body))

	// Verifica se o status é 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("falha ao obter informações do usuário: status %d, resposta: %s", resp.StatusCode, string(body))
	}

	// Processa o corpo da resposta como JSON
	var userInfo map[string]interface{}
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, fmt.Errorf("erro ao parsear resposta JSON: %w", err)
	}

	return userInfo, nil

}

// GetCities obtém a lista de cidades com paginação
func (c *FortlevClient) GetCities(currentPage, docsPerPage int) (*CityResponse, error) {
	url := fmt.Sprintf("%s/partner/brazilian-city/state?current_page=%d&docs_per_page=%d",
		c.BaseURL, currentPage, docsPerPage)

	// Cria a requisição HTTP GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	// Adiciona os headers necessários
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.AuthToken)

	// Executa a requisição
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %w", err)
	}

	// Verifica o status da resposta
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na requisição: status %d, resposta: %s", resp.StatusCode, string(body))
	}

	// Decodifica a resposta JSON
	var cityResponse CityResponse
	if err := json.Unmarshal(body, &cityResponse); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	return &cityResponse, nil
}
