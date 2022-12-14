# password

# Sobre o projeto

password é uma api construída para ter rotas que auxiliam sobre senhas.

A API atualmente tem uma única rota que verifica se a senha é forte ou não.

# Como executar o projeto

## Com o make

```bash
# Instalar dependencias do projeto
make install

# Executar o projeto
make run
```
### testes unitários

```bash
# Rodar testes
make test

# Rodar testes e abrir aba para visualização da cobertura de testes
make coverage
```

## Sem o make

```bash
# Instalar dependencias do projeto
go get
go mod tidy
go mod vendor

# Executar o projeto
go run main.go
```
### testes unitários

```bash
# Rodar testes
go test ./...
```

# GraphQL

Inicie o projeto, url base: http://localhost:8080/graphql

# Rest

Descomente a "config.NewRestServer()" no arquivo main.go e comente a linha config.NewGraphQLServer()

Inicie o projeto, url base: http://localhost:8080
