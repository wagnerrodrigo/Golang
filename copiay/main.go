package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/atotto/clipboard"
)

func findLatestFile(dir string) (string, error) {
	var latestFile string
	var latestTime time.Time

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".mp4" { // Exemplo para arquivos MP4
			fileTime := file.ModTime()
			if fileTime.After(latestTime) {
				latestTime = fileTime
				latestFile = filepath.Join(dir, file.Name())
			}
		}
	}

	return latestFile, nil
}

func main() {
	dir := "/home/wagner/Downloads/teste"
	latestFile, err := findLatestFile(dir)
	if err != nil {
		fmt.Println("Erro ao buscar o arquivo:", err)
		return
	}

	/// FUNCIONOU MAS NÃO DEU CERTO PARA O MEU MOTIVO

	fmt.Println("Arquivo mais recente:", latestFile)
	// Aqui você poderia copiar o arquivo para a área de transferência.
	clipboard.WriteAll(latestFile)

	// latestFile, err := findLatestFile(dir)
	// if err != nil {
	// 	fmt.Println("Erro ao buscar o arquivo:", err)
	// 	return
	// }

	// fmt.Println("Arquivo mais recente:", latestFile)

	// Copiar o caminho do arquivo para a área de transferência (Linux exemplo)
	// exec.Command("xdotool", "search", "--onlyvisible", "--class", "firefox", "windowactivate").Run()
	// time.Sleep(2000 * time.Millisecond) // Pequeno delay para garantir que a janela está ativa
	// exec.Command("xdotool", "key", "ctrl+v").Run()
}

// package main

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/chromedp/chromedp"
// )

// func main() {
// 	ctx, cancel := chromedp.NewContext(context.Background())
// 	defer cancel()

// 	if err := chromedp.Run(ctx,
// 		chromedp.Navigate("https://web.whatsapp.com"),
// 		chromedp.Sleep(5*time.Second), // Tempo para autenticar o WhatsApp
// 		chromedp.Click(`#app > div > div > div:nth-child(3) > div > div`, chromedp.NodeVisible),
// 		chromedp.SendKeys(`#input-file`, "/home/wagner/Downloads/preenceher_chamado.mp4"),
// 	); err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Vídeo enviado com sucesso!")
// }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"time"
// )

// func main() {
// 	videoPath := os.Getenv("HOME") + "/Downloads/preenceher_chamado.mp4"

// 	// Verifique se o arquivo existe
// 	if _, err := os.Stat(videoPath); os.IsNotExist(err) {
// 		fmt.Println("O arquivo de vídeo não foi encontrado em", videoPath)
// 		return
// 	}

// 	// Copie o vídeo para a área de transferência usando xclip
// 	copyCmd := exec.Command("xclip", "-selection", "clipboard", "-t", "video/mp4", "-i", videoPath)
// 	if err := copyCmd.Run(); err != nil {
// 		fmt.Println("Erro ao copiar o vídeo para a área de transferência:", err)
// 		return
// 	}

// 	// Use xdotool para encontrar a janela do WhatsApp Web e focá-la
// 	focusCmd := exec.Command("xdotool", "search", "--onlyvisible", "--name", "WhatsApp")
// 	output, err := focusCmd.Output()
// 	if err != nil || len(output) == 0 {
// 		fmt.Println("Erro ao localizar a janela do WhatsApp Web ou janela não encontrada.")
// 		return
// 	}

// 	// Foca na janela do WhatsApp Web (usa o primeiro ID retornado)
// 	windowID := string(output[:len(output)-1]) // Remove a nova linha
// 	exec.Command("xdotool", "windowactivate", windowID).Run()

// 	// Emule o atalho de teclado Ctrl+V para colar o vídeo
// 	time.Sleep(1 * time.Second) // Pequeno delay para garantir que a janela esteja em foco
// 	pasteCmd := exec.Command("xdotool", "key", "--delay", "500", "ctrl+v")
// 	if err := pasteCmd.Run(); err != nil {
// 		fmt.Println("Erro ao colar o vídeo:", err)
// 		return
// 	}

// 	// Pressione Enter para enviar o vídeo
// 	time.Sleep(8 * time.Second)
// 	enterCmd := exec.Command("xdotool", "key", "Return")
// 	if err := enterCmd.Run(); err != nil {
// 		fmt.Println("Erro ao enviar o vídeo:", err)
// 		return
// 	}

// 	fmt.Println("Vídeo enviado com sucesso!")
// }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"time"
// )

// func main() {
// 	videoPaht := os.Getenv("HOME") + "/Downloads/preenceher_chamado.mp4"
// 	// fmt.Println(videoPaht)

// 	if _, err := os.Stat(videoPaht); os.IsNotExist(err) {
// 		fmt.Println("O arquivo de video não encontado em ", videoPaht)
// 		return
// 	}

// 	copyCmd := exec.Command("xclip", "-selection", "clipboard", "-t", "video/mp4", "-i", videoPaht)
// 	if err := copyCmd.Run(); err != nil {
// 		fmt.Println("Error ao copiar o video para a area de transferencia", err)
// 		return
// 	}

// 	pastCmd := exec.Command("xdotool", "key", "--delay", "500", "ctrl+v")
// 	if err := pastCmd.Run(); err != nil {
// 		fmt.Println("Error ao colar o video:", err)
// 		return
// 	}

// 	// openCmd := exec.Command("xdg-open", "https://web.whatsapp.com")
// 	// if err := openCmd.Start(); err != nil {
// 	// 	fmt.Println("Erro ao abrir o WhatsApp Web:", err)
// 	// 	return
// 	// }

// 	// Use xdotool para encontrar a janela do WhatsApp Web e focá-la
// 	focusCmd := exec.Command("xdotool", "search", "--onlyvisible", "--name", "WhatsApp")
// 	output, err := focusCmd.Output()
// 	if err != nil || len(output) == 0 {
// 		fmt.Println("Erro ao localizar a janela do WhatsApp Web ou janela não encontrada.")
// 		return
// 	}

// 	// Foca na janela do WhatsApp Web (usa o primeiro ID retornado)
// 	windowID := string(output[:len(output)-1]) // Remove a nova linha
// 	exec.Command("xdotool", "windowactivate", windowID).Run()

// 	time.Sleep(5 * time.Second)
// 	entreCmd := exec.Command("xdotool", "key", "Return")
// 	if err := entreCmd.Run(); err != nil {
// 		fmt.Println("Error ao enviar o video", err)
// 		return
// 	}
// 	fmt.Println("video enviado com sucesso ")
// }
