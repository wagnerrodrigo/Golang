# SendMail
Uso da biblioteca do go para envio de [email](https://pkg.go.dev/gopkg.in/gomail.v2#example-package). 

## Como usar
-  Configure
  
        host := 
        port := 587
	    user := ""
	    password := ""
> Utilizei o [serviço do MailTrap](https://mailtrap.io/) para configurar o host, port, user e password 


- Uma vez configurando as informações acima basta rodar o comando.
  
~~~go 
go rum email.go
~~~
> Pronto sua mensagem será exibida na caixa de e-mail do MailTrap.