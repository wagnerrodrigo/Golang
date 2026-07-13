package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// 1. Configurar o Chromedp
	// Por padrão, o chromedp roda em modo "headless" (background).
	// Para ver a automação acontecendo no seu monitor, vamos desabilitar o headless (headless = false).
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // Mude para true para rodar em background
		chromedp.Flag("disable-gpu", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Criar o contexto do chromedp
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// Definir um tempo limite (timeout) para o script não travar infinitamente
	// caso um botão ou elemento demore muito para aparecer
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Variáveis para guardar informações extraídas, se necessário
	var pageTitle string
	var screenshotBuf []byte

	// URL da Intranet
	targetURL := "http://intranet.exemplo.com.br"

	log.Printf("Iniciando navegação para: %s", targetURL)

	// 2. Executar as ações
	err := chromedp.Run(ctx,
		// Navegar até a URL
		chromedp.Navigate(targetURL),

		// Esperar o corpo da página carregar (ou um elemento específico)
		// Exemplo: chromedp.WaitVisible(`#id_do_elemento`, chromedp.ByID)
		chromedp.WaitVisible(`body`, chromedp.ByQuery),

		// Capturar o título da página
		chromedp.Title(&pageTitle),

		// =========================================================================
		// EXEMPLOS DE INTERAÇÕES (SUBSTITUA PELOS SELETORES DA SUA PÁGINA)
		// =========================================================================

		// Exemplo 1: Clicar em um botão por ID
		chromedp.Click(`#botao-iniciar`, chromedp.ByID),

		// Exemplo 2: Digitar em um input
		// chromedp.SendKeys(`#input-cpf`, "123.456.789-00", chromedp.ByID),

		// Exemplo 3: Clicar em um botão pelo seletor CSS (ex: classe ou tag)
		// chromedp.Click(`button.btn-submit`, chromedp.ByQuery),

		// Exemplo 4: Esperar 2 segundos para dar tempo de ver a ação ou a transição
		chromedp.Sleep(2*time.Second),

		// Exemplo 5: Tirar um screenshot da tela final
		chromedp.FullScreenshot(&screenshotBuf, 90),
	)

	if err != nil {
		log.Fatalf("Erro durante a execução do programa: %v", err)
	}

	log.Printf("Página acessada com sucesso! Título: %s", pageTitle)

	// Salvar o screenshot no disco
	if len(screenshotBuf) > 0 {
		screenshotPath := "screenshot.png"
		err = ioutil.WriteFile(screenshotPath, screenshotBuf, 0o644)
		if err != nil {
			log.Printf("Aviso: Não foi possível salvar o screenshot: %v", err)
		} else {
			log.Printf("Screenshot salvo como: %s", screenshotPath)
		}
	}
}
