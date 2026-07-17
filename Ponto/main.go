package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Não foi possível carregar o arquivo .env, usando variáveis do sistema se existirem")
	}

	// Obter a URL da variável de ambiente
	targetURL := os.Getenv("URL")
	if targetURL == "" {
		log.Fatal("Erro: A variável 'URL' não está configurada no arquivo .env ou no sistema")
	}

	targetMatricula := os.Getenv("MATRICULA")
	if targetMatricula == "" {
		log.Fatal("Erro: A variável 'MATRICULA' não está configurada no arquivo .env ou no sistema")
	}

	targetPassword := os.Getenv("PASSWORD")
	if targetPassword == "" {
		log.Fatal("Erro: A variável 'PASSWORD' não está configurada no arquivo .env ou no sistema")
	}

	// 1. Configurar o Chromedp
	// Por padrão, o chromedp roda em modo "headless" (background).
	// Para ver a automação acontecendo no seu monitor, vamos desabilitar o headless (headless = false).
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true), // Mude para true para rodar em background
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true), // Necessário para rodar em alguns ambientes, como o GitHub Actions
		chromedp.Flag("disable-setuid-sandbox", true),
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

	log.Printf("Iniciando navegação para: %s", targetURL)

	// 2. Executar o login
	err = chromedp.Run(ctx,
		// Navegar até a URL
		chromedp.Navigate(targetURL),

		// Esperar o corpo da página carregar
		chromedp.WaitVisible(`body`, chromedp.ByQuery),

		// Capturar o título da página
		chromedp.Title(&pageTitle),

		// Esperar o botão "Iniciar Pesquisa" estar visível e clicar
		chromedp.WaitVisible(`button.pesquisar`, chromedp.ByQuery),
		chromedp.Click(`button.pesquisar`, chromedp.ByQuery),

		// Preencher matrícula e senha
		chromedp.WaitVisible(`#matricula`, chromedp.ByID),
		chromedp.SendKeys(`#matricula`, targetMatricula, chromedp.ByID),

		chromedp.WaitVisible(`#password`, chromedp.ByID),
		chromedp.SendKeys(`#password`, targetPassword, chromedp.ByID),

		// Clicar em Login
		chromedp.WaitVisible(`a.iniciar`, chromedp.ByQuery),
		chromedp.Click(`a.iniciar`, chromedp.ByQuery),

		// Aguardar 3 segundos para carregar a resposta do login (o telefone ou o SweetAlert)
		chromedp.Sleep(3*time.Second),
	)
	if err != nil {
		log.Fatalf("Erro durante o login: %v", err)
	}

	// 3. Verificar se a pesquisa já foi realizada hoje
	var isAlreadyDone bool
	err = chromedp.Run(ctx,
		chromedp.Evaluate(`document.querySelector('.swal-modal') !== null`, &isAlreadyDone),
	)
	if err != nil {
		log.Printf("Aviso: Falha ao verificar se pesquisa já estava concluída: %v", err)
	}

	if isAlreadyDone {
		log.Println("A pesquisa de saúde já foi realizada hoje! Encerrando com sucesso...")

		// Clicar no botão "OK" do SweetAlert e tirar screenshot da tela de aviso
		err = chromedp.Run(ctx,
			chromedp.FullScreenshot(&screenshotBuf, 90),
			chromedp.Sleep(1*time.Second), // validar a ordem de execução para garantir que o screenshot seja tirado antes do clique
			chromedp.Click(`button.swal-button--confirm`, chromedp.ByQuery),
		)
		if err != nil {
			log.Printf("Aviso: Não foi possível clicar no OK do SweetAlert: %v", err)
		}
	} else {
		log.Println("Pesquisa ainda não realizada. Iniciando preenchimento da pesquisa...")

		// 4. Executar os passos da pesquisa
		err = chromedp.Run(ctx,
			// Esperar a próxima tela carregar (confirmar contato)
			chromedp.WaitVisible(`button[onclick="telefone()"]`, chromedp.ByQuery),
			chromedp.Click(`button[onclick="telefone()"]`, chromedp.ByQuery),

			// AGUARDAR A ANIMAÇÃO (slideDown de 1 segundo) TERMINAR COMPLETAMENTE
			chromedp.Sleep(1500*time.Millisecond),

			// Esperar o botão "ESTOU BEM" carregar e clicar nele
			chromedp.WaitVisible(`#pergunta4 button.btn-success`, chromedp.ByQuery),
			chromedp.Click(`#pergunta4 button.btn-success`, chromedp.ByQuery),

			// AGUARDAR O MODAL DO ALERTIFY APARECER E COMPLETAR A ANIMAÇÃO
			chromedp.Sleep(1*time.Second),

			// Esperar o botão "Concluir" do modal do Alertify aparecer e clicar nele
			chromedp.WaitVisible(`button.ajs-ok`, chromedp.ByQuery),
			chromedp.Click(`button.ajs-ok`, chromedp.ByQuery),

			// Esperar 4 segundos (o AJAX do site leva 3 segundos para redirecionar após salvar)
			chromedp.Sleep(4*time.Second),

			// Tirar um screenshot da tela final
			chromedp.FullScreenshot(&screenshotBuf, 90),
		)
		if err != nil {
			log.Fatalf("Erro durante o preenchimento da pesquisa: %v", err)
		}
	}

	log.Printf("Fluxo concluído com sucesso! Título da página: %s", pageTitle)

	// Salvar o screenshot no disco
	if len(screenshotBuf) > 0 {
		screenshotPath := "screenshot.png"
		if isAlreadyDone {
			screenshotPath = "screenshot_already_done.png"
			log.Println("Salvando screenshot da tela de aviso de pesquisa já realizada...")
		}
		err = ioutil.WriteFile(screenshotPath, screenshotBuf, 0o644)
		if err != nil {
			log.Printf("Aviso: Não foi possível salvar o screenshot: %v", err)
		} else {
			log.Printf("Screenshot salvo como: %s", screenshotPath)
		}
	}
}
