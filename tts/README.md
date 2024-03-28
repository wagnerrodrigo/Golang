## TTS

Implementação simples de Text-To-Speech (texto para fala).


como usar 

Baixar a biblioteca.
```go 
go get github.com/hegedustibor/htgo-tts
```

Em seguida, abra o código tts.go e na linha de número 16, mude o texto entre aspas para o texto desejado.

```go 
speech.Speak("Digite o texto aqui entre as aspas")
```

em seguida execute o comando 

```go 
 go run tts.go
```

Exemplo
```go 
 go run tts.go
```
> saida esperada: um áudio com a frase --> **Esse texto vai virar um áudio.**

uma pasta chamada áudio e dentro o texto convertido em um arquivo mp3.



