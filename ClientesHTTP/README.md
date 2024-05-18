# clientesHTTP


Implementação de clientes http

cliente criados 

- get-user
- login
- get-current-user


como usar 

Para buscar usuário digite o comando abaixo.
```go 
    go run ./cmd/client -action get-user -id 1                              
```
> saida esperada:  {1 atuny0 9uQFF1Lh atuny0@sohu.com }


Em seguida, com os dados de usuário e login digite o comando abaixo.
```go 
    go run ./cmd/client -action login -user atuny0 -pass 9uQFF1Lh 
```
> saida esperada: 
> 
> usuario logado com sucesso
> 
> {1 atuny0  atuny0@sohu.com eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJhdHVueTAiLCJlbWFpbCI6ImF0dW55MEBzb2h1LmNvbSIsImZpcnN0TmFtZSI6IlRlcnJ5IiwibGFzdE5hbWUiOiJNZWRodXJzdCIsImdlbmRlciI6Im1hbGUiLCJpbWFnZSI6Imh0dHBzOi8vcm9ib2hhc2gub3JnL1RlcnJ5LnBuZz9zZXQ9c2V0NCIsImlhdCI6MTcxNjA1MzM3NCwiZXhwIjoxNzE2MDU2OTc0fQ.jMZQAIaKjlsi629ZSBCSdOe5-VielinA1Rh1ZkJmwfc} 


E por último o cliente customizado.

```go 
    go run ./cmd/client -action get-current-user -user atuny0 -pass 9uQFF1Lh
```
> saida esperada: {1 atuny0 9uQFF1Lh atuny0@sohu.com }