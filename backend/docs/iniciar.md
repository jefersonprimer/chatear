# Guia de Início Rápido para o Projeto Chatear-Backend

Este guia detalha como configurar, executar e testar o projeto Chatear-Backend.

## 1. Pré-requisitos

Certifique-se de ter as seguintes ferramentas instaladas em sua máquina:

*   **Go:** Versão 1.21 ou superior.
*   **Docker:** Para executar os serviços de banco de dados, cache e mensageria.
*   **Docker Compose:** Para orquestrar os contêineres Docker.
*   **[golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate):** Para executar as migrações do banco de dados.
*   **Insomnia (ou Postman):** Para testar as APIs HTTP e GraphQL.

## 2. Configuração do Projeto

1.  **Clonar o Repositório:**
    ```bash
    git clone https://github.com/jefersonprimer/chatear-backend.git
    cd chatear-backend
    ```

2.  **Instalar Dependências Go:**
    ```bash
    go mod download
    ```

3.  **Configurar Variáveis de Ambiente:**
    Crie um arquivo `.env` na raiz do projeto, copiando o conteúdo de `env.example` e preenchendo as variáveis conforme necessário.

    ```bash
    cp env.example .env
    ```

    **É crucial que você preencha o `SUPABASE_CONNECTION_STRING` com a sua connection string do Supabase, e os outros segredos como `JWT_SECRET` e as configurações de SMTP.**

## 3. Executando a Aplicação

Você pode executar a aplicação de duas maneiras: usando Docker Compose (recomendado para um ambiente semelhante à produção) ou localmente para desenvolvimento e depuração.

### 3.1. Usando Docker Compose (Recomendado)

A maneira mais simples de executar todos os componentes da aplicação (API, workers, NATS, Redis) é usando o `docker-compose`.

1.  **Iniciar os Contêineres:**
    ```bash
    docker-compose -f docker-compose.events.yml up --build
    ```
    Isso irá construir as imagens e iniciar todos os serviços.

2.  **Verificar o Status:**
    ```bash
    docker-compose -f docker-compose.events.yml ps
    ```

A API estará disponível em `http://localhost:8080`.

### 3.2. Localmente (para Desenvolvimento)

Para desenvolvimento e depuração, você pode executar a API e os workers localmente usando os comandos do `Makefile`.

1.  **Executar a API:**
    ```bash
    make run-api
    ```

2.  **Executar os Workers (em terminais separados):**
    ```bash
    make run-worker-notification
    make run-worker-user-delete
    make run-worker-user-hard-delete
    make run-worker-user-permanent-deletion-scheduler
    make run-worker-user-registered
    ```

## 4. Migrações do Banco de Dados

Antes de iniciar a aplicação, você precisa aplicar as migrações do banco de dados no seu banco de dados Supabase.

**Nota:** A configuração do Docker Compose não inclui um banco de dados. Você deve usar um banco de dados externo, como o Supabase.

Execute o seguinte comando, substituindo a sua connection string do Supabase:

```bash
migrate -database "SUA_SUPABASE_CONNECTION_STRING" -path migrations/postgres up
```

## 5. Testando a Aplicação

### 5.1. Testes Automatizados

Para executar os testes automatizados, use o seguinte comando:

```bash
make test
```

### 5.2. Testando com Insomnia

Você pode testar as APIs HTTP e GraphQL usando o Insomnia. O endpoint da API GraphQL é `http://localhost:8080/graphql`.

Consulte o restante deste documento para obter exemplos de queries e mutations.

## 6. Comandos Úteis do `Makefile`

O `Makefile` contém vários comandos úteis para o desenvolvimento:

*   `make build`: Compila a API e os workers.
*   `make run-api`: Executa a API localmente.
*   `make run-worker-*`: Executa um worker específico localmente.
*   `make test`: Executa os testes automatizados.
*   `make lint`: Executa o `go vet` para análise estática.
*   `make clean`: Remove os binários compilados.
