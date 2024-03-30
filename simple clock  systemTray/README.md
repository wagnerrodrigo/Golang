## Simple relogio no systray

Implementação simples de relogio no system tray.


como usar 

Excute o binario corresponde ao seu sistema operaciona que se encontra na pasta **_applications_** 

> Obs1: o main.exe para o win, ja o main para linux

>*Obs2: os executáveis que estão na pasta applications foram criado a partir do sistema operacional linux, caso não funcione compile localmente, seguindo os passos abaixo.

Baixar a biblioteca.
```go 
go get github.com/Osuka42g/simple-clock-systray
```


em seguida execute o comando, em quando o código esta executando no terminal a app estará no system tray

```go 
 go run main.go   
```
para gerar um executavel 

```go 
 go build main.go
```
o arquivo para executar estará na mesma pasta onde o projeto foi baixado e executado.

Exemplo
```go 
 go build main.go
```
> saida esperada: um arquivo executavel com o nome main estara na pasta dos arquivos





