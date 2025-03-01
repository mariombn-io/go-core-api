#!/bin/bash

echo "🛠 Criando estrutura do projeto Go ..."

# Criar a estrutura de diretórios
mkdir -p cmd/cli api/{middlewares,requests,handlers,resources} entities \
         services/{exemple,user} repositories mappers helpers logger config \
         policies tests docs

# Criar arquivos principais
touch main.go .env go.mod go.sum

# Criar arquivos de configuração
touch config/config.go
touch logger/logger.go

# Criar arquivos da API
touch api/api.go api/route.go api/server.go
touch api/middlewares/jwt.go
touch api/requests/exemple.go
touch api/handlers/exemple.go
touch api/resources/exemple.go

# Criar arquivos das entidades
touch entities/exemple.go entities/user.go

# Criar serviços
touch services/exemple/create.go
touch services/user/create.go

# Criar repositórios
touch repositories/exemple.go
touch repositories/user.go

# Criar mapeadores
touch mappers/exemple_mapper.go
touch mappers/user_mapper.go

# Criar helpers
touch helpers/formatter.go helpers/valid.go helpers/cache.go

# Criar políticas
touch policies/user.go

# Criar estrutura da CLI
touch cmd/cli/cli.go

# Criar um teste de integração inicial
touch tests/api_middleware_test.go

# Criar README e documentação inicial
touch docs/README.md

# Mensagem de sucesso
echo "✅ Estrutura do projeto criada com sucesso!"
echo "📂 Arquivos e diretórios configurados."
