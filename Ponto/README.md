# Guia de Automação de Navegador em Go (com `chromedp`)

Este projeto contém um script em Go estruturado para acessar a página `https://intranet.rivelli.ind.br/app/pesquisa_saude/index.php` e interagir com botões e formulários.

## 🛠️ Tecnologias
- **Go (Golang)**
- **chromedp**: Biblioteca para controle do Chrome/Chromium via Chrome DevTools Protocol (CDP).

## 🚀 Como Rodar o Script

1. Certifique-se de que o **Google Chrome** ou **Chromium** está instalado no seu sistema.
2. Na pasta do projeto, execute o comando:
   ```bash
   go run main.go
   ```

## 🔍 Como Ajustar os Cliques (Seletores CSS)

No arquivo `main.go`, você precisará especificar os seletores dos elementos que deseja clicar ou preencher. Para encontrar estes seletores:
1. Abra o Chrome na página desejada.
2. Clique com o botão direito no botão/campo e selecione **Inspecionar**.
3. Veja o `id`, `class` ou outros atributos do elemento HTML.

### Exemplos comuns no `main.go`:
- **Clicar por ID:** `chromedp.Click("#id-do-botao", chromedp.ByID)`
- **Clicar por classe CSS:** `chromedp.Click(".classe-do-botao", chromedp.ByQuery)`
- **Digitar texto:** `chromedp.SendKeys("#id-do-input", "seu texto", chromedp.ByID)`
- **Esperar elemento carregar:** `chromedp.WaitVisible("#id-do-elemento", chromedp.ByID)`

## 👁️ Modo Visual vs. Background (Headless)
No arquivo `main.go`, o navegador está configurado para abrir visualmente (`headless = false`). Se você quiser rodar a automação silenciosamente em background, altere a flag no código:
```go
chromedp.Flag("headless", true)
```
