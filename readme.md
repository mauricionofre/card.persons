# Person API

Person API é uma aplicação CRUD (Create, Read, Update, Delete) para gerenciar informações de pessoas, utilizando Go, PostgreSQL e RabbitMQ.

## Características

- API RESTful para operações CRUD de pessoas
- Integração com PostgreSQL para armazenamento de dados
- Uso de RabbitMQ para publicação de eventos
- Estrutura de projeto seguindo as melhores práticas de Go

## Pré-requisitos

- Go 1.16+
- PostgreSQL
- RabbitMQ
- Docker (opcional)

## Configuração

1. Clone o repositório:
   ```
   git clone https://github.com/seu-usuario/person-api.git
   cd person-api
   ```

2. Copie o arquivo `.env.example` para `.env` e preencha com suas configurações:
   ```
   cp .env.example .env
   ```

3. Edite o arquivo `.env` com suas configurações de banco de dados e RabbitMQ.

4. Instale as dependências:
   ```
   go mod tidy
   ```

5. Crie a tabela no banco de dados PostgreSQL:
   ```sql
   CREATE TABLE persons (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       age INT NOT NULL
   );
   ```

## Executando a aplicação

Para iniciar a aplicação, execute:

```
go run cmd/api/main.go
```

A API estará disponível em `http://localhost:8080`.

## Endpoints

- `POST /persons`: Criar uma nova pessoa
- `GET /persons/:id`: Obter uma pessoa por ID
- `PUT /persons/:id`: Atualizar uma pessoa existente
- `DELETE /persons/:id`: Excluir uma pessoa

## Estrutura do Projeto

```
.
├── cmd
│   └── api
│       └── main.go
├── internal
│   ├── config
│   ├── handler
│   ├── model
│   ├── repository
│   └── service
├── pkg
│   └── rabbitmq
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Eventos RabbitMQ

A aplicação publica os seguintes eventos na exchange "person":

- `person.created`: Quando uma nova pessoa é criada
- `person.updated`: Quando uma pessoa é atualizada
- `person.deleted`: Quando uma pessoa é excluída

## Contribuindo

Contribuições são bem-vindas! Por favor, sinta-se à vontade para submeter um Pull Request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
