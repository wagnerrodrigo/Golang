## Maps

estudo sobre map.

como usar 

> forma de execurar o programa em go  **{go run}** **{nome_do_aqruivo}**


```go 
go run maps.go {Insira o texto }
```

Exemplo 1


```go 
go run maps.go \Lorem ipsum dolor sit amet leo eu velit ante sagittis dolor \ turpis dis
```


```
saida esperada >>

E = 1
V = 1
T = 1
L = 2
I = 1
D = 3
S = 2
A = 2
```

Exemplo 2

```go 
go run maps.go O Lorem Ipsum é um texto modelo da indústria tipográfica e de impressão. O Lorem Ipsum tem vindo a ser o texto padrão usado por estas indústrias desde o ano de 1500, quando uma misturou os caracteres de um texto para criar um espécime de livro. Este texto não só sobreviveu 5 séculos, mas também o salto para a tipografia electrónica, mantendo-se essencialmente inalterada. Foi popularizada nos anos 60 com a disponibilização das folhas de Letraset, que continham passagens com Lorem Ipsum, e mais recentemente com os programas de publicação como o Aldus PageMaker que incluem versões do Lorem Ipsum.

```

```
saida esperada >>

E = 7
V = 2
N = 2
6 = 1
A = 6
P = 9
R = 1
O = 8
I = 9
M = 5
D = 11
5 = 1
F = 2
L = 6
T = 8
1 = 1
Q = 3
Ã = 1
U = 5
S = 5
C = 7

```

**[Link para gerar texto](https://pt.lipsum.com/)**



