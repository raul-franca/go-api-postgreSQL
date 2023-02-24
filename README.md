# API Todo List em Go

Esta é uma API de lista de tarefas (Todo List) em Go que utiliza um banco de dados PostgreSQL. O objetivo deste projeto é fornecer um exemplo de como implementar uma API RESTful em Go com um banco de dados PostgreSQL, utilizando técnicas para gerar código de forma mais rápida e com menos erros.

## Tecnologias utilizadas

-   [Go](https://golang.org/)
-   [PostgreSQL](https://www.postgresql.org/)
-   [SQLC](https://sqlc.dev/)
-   [Golang-migrate](https://github.com/golang-migrate/migrate)
-   [Swagger](https://swagger.io/)

## Instalação

1.  Clone este repositório

    bashCopy code

    `git clone https://github.com/raul-franca/go-api-postgreSQL.git`


2.  Execute o comando abaixo para baixar as dependências

    goCopy code

    `go mod download`

3.  Crie o banco de dados PostgreSQL com o nome `todo_list_db`
4.  Crie as tabelas executando o arquivo `db/migrations/migrate.go`

    goCopy code

    `go run db/migrations/migrate.go`

5.  Execute a aplicação

    goCopy code

    `go run main.go`

6.  Acesse a documentação da API em `http://localhost:8080/swagger/index.html`

## Endpoints

### GET /tasks

Retorna todas as tarefas cadastradas.

### GET /tasks/{id}

Retorna uma tarefa específica pelo seu ID.

### POST /tasks

Cadastra uma nova tarefa.

### PUT /tasks/{id}

Atualiza uma tarefa existente pelo seu ID.

### DELETE /tasks/{id}

Remove uma tarefa existente pelo seu ID.

## Documentação da API

A documentação da API foi gerada automaticamente utilizando a ferramenta Swagger. Para acessá-la, execute a aplicação e acesse `http://localhost:8080/swagger/index.html` em seu navegador.

## Autor

raulmfranca@gmail.com

## Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo LICENSE para obter mais informações.