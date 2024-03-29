## Copy file

Implementação simples de copia de arquivos.


como usar 


1° baixe o código git clone https://github.com/wagnerrodrigo/Golang.git 
abra a pasta do projeto CopiaArquivo

2° crie um arquivo e edit.

3° Abra o código copy.go e na linha de número 10, mude o texto entre aspas para o **nome do arquivo criado anteriormente**.

```go
EXEMPLO -->> os.Open("arquivo_original.txt")
```

4° Na linha de número 16, mude o texto entre aspas para o nome do **arquivo de saida**

```go
EXEMPLO -->> os.Open("arquivo_destino.txt")
```


em seguida execute o comando 

```go 
 go run copy.go
```

> saida esperada: um arquivo será criado no mesmo diretorio com o nome e conteúdo salvo no primeiro arquivo





