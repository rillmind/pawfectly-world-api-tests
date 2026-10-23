# pawfectly-world-api-tests

Testes de integração, em Go, da API da [Pawfectly World](https://github.com/rillmind/pawfectly-world-api).

## O que faz

Faz requisições HTTP reais para a API rodando e confere os status de resposta. O `TestMain`
(`test/e2e/main_test.go`) checa `GET /health` e faz login como administrador; depois rodam os
testes de usuário (`test/e2e/user_test.go`):

- `TestCreateUser`: `POST /user` devolve 201;
- `TestListUsers`: `GET /user` devolve 200;
- `TestListUserByID`, `TestUpdateUserByID`, `TestDeleteUserByID`: buscam, atualizam e removem
  um usuário criado só para o teste.

Cada teste cria e apaga os próprios dados, então pode rodar sozinho (`go test -run`). Os testes
de autenticação (`auth_test.go`) ainda são um TODO.

## Por que existe

A API da Pawfectly World foi meu TCC do técnico (2023) e só tinha os testes de exemplo do template
do Nest. Em 2026 voltei a ela para cobrir os endpoints com testes de integração escritos em Go,
só com a biblioteca padrão (`testing` + `net/http`).

## Instalação

Pré-requisitos: Go 1.26+ e a API da Pawfectly World rodando (veja o README dela).

```sh
git clone https://github.com/rillmind/pawfectly-world-api-tests.git
cd pawfectly-world-api-tests
go mod download
```

Crie um `.env` na raiz (já está no `.gitignore`):

```env
API_BASE_URL=http://127.0.0.1:2306
ADM_USERNAME=admin
ADM_PASSWORD=123456
```

O usuário precisa ter o papel `admin`. A API cria todo usuário como `user`, então promova-o
direto no banco. Com o MongoDB do compose da API:

```sh
podman exec pawfectly-mongodb mongosh --quiet pawfectly_world \
  --eval 'db.users.updateOne({username: "admin"}, {$set: {role: "admin"}})'
```

## Como testar

```sh
go vet ./...
go test -count=1 ./...
```

Se a API não responder 200 em `GET /health`, o `TestMain` encerra a suíte antes de qualquer
teste (`status-code: 404, want 200`). Hoje a versão publicada da API ainda não tem essa rota,
então a suíte para aí. O `go vet ./...` passa.
