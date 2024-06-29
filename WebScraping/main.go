package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// res, err := http.Get("https://www.google.com/")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer res.Body.Close()

	// input := res.Body

	input, _ := os.Open("index.html")

	// Primeira opção para pegar os token
	// tokenizer := html.NewTokenizer(input)

	// for {
	// 	tokenType := tokenizer.Next()
	// 	token := tokenizer.Token()

	// 	if tokenType == html.ErrorToken {
	// 		// chegou ao fim do arquivo
	// 		if tokenizer.Err() == io.EOF {
	// 			return
	// 		}
	// 		fmt.Println(tokenizer.Err())
	// 	}

	// 	// CommentToken pegar comentarios
	// 	// if tokenType == html.CommentToken {
	// 	// 	fmt.Println(token)
	// 	// }

	// 	// self close tag fechada em si mesmo SelfClosingTagToken exemplo </br>
	// 	// if tokenType == html.SelfClosingTagToken {
	// 	// 	fmt.Println(token)
	// 	// }

	// 	// imprimir todos os links da pagina
	// 	if tokenType == html.StartTagToken {
	// 		if token.Data == "a" {
	// 			for _, att := range token.Attr {

	// 				// fmt.Println(att.Key, att.Val)
	// 				fmt.Println(att.Val)
	// 			}
	// 			tt := tokenizer.Next()
	// 			if tt == html.TextToken {
	// 				fmt.Println(tokenizer.Token())
	// 			}
	// 		}
	// 	}

	// }

	// # Segunda opção
	// node, _ := html.Parse(input)

	// var acessa func(*html.Node)

	// // func´~ao recursiva
	// acessa = func(n *html.Node) {

	// 	// fmt.Println(n.Data)

	// 	// pegando os link
	// 	if n.Type == html.ElementNode && n.Data == "a" {
	// 		// fmt.Println(n) pega com os bit
	// 		fmt.Print(n.FirstChild.Data + " -> ")
	// 		for _, a := range n.Attr {
	// 			if a.Key == "href" {
	// 				fmt.Println(a.Val)
	// 			}
	// 		}

	// 	}

	// 	// entra no conteudo e pega o proximo conteudo
	// 	for next := n.FirstChild; next != nil; next = next.NextSibling {
	// 		acessa(next)
	// 	}
	// }

	// acessa(node)

	// # terceira forma de WebScraping
	doc, _ := goquery.NewDocumentFromReader(input)

	doc.Find(".foo").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}
