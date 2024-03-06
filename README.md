# Golang RESTful API

Esta é uma aplicação RESTful simples escrita em Go, que utiliza um banco de dados PostgreSQL para armazenar informações sobre usuários.

## Pré-requisitos

- Go 1.16 ou superior
- PostgreSQL
- Git

## Configuração

1. Clone o repositório:

```bash
git clone https://github.com/fmedeiros95/golang-rest-api.git
```


2. Instale as dependências:

```bash
cd golang-rest-api
go mod tidy
```


3. Configure as variáveis de ambiente:

Crie um arquivo `.env` na raiz do projeto e defina as variáveis de ambiente necessárias. Você pode usar o arquivo `.env.example` como referência.

4. Inicie o servidor:
```bash
go run main.go
```


## Rotas

### Usuários

- `GET /users`: Retorna uma lista de todos os usuários.
- `GET /users/{id}`: Retorna os detalhes de um usuário específico.
- `POST /users`: Cria um novo usuário.
- `PUT /users/{id}`: Atualiza os detalhes de um usuário existente.
- `DELETE /users/{id}`: Exclui um usuário existente.


## Contribuição

Sinta-se à vontade para enviar pull requests, relatar problemas ou fornecer sugestões para este projeto.


## Licença

Este projeto está licenciado sob a [Licença MIT](https://opensource.org/licenses/MIT).