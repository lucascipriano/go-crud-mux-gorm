- Criar rotas
- Criar User model
- Adicionar goORM

Ideias futuras
- bcrypt para senhas
- gerar JWT


## docker

Comando utilizado para criar a database
`docker run -d --name api-golang -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=apigolang -e MYSQL_USER=golang -e MYSQL_PASSWORD=golang mysql:latest`
