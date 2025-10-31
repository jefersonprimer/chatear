### Como usar
Ambiente de desenvolvimento (mantém shell e debug):
  docker build -t chatear-backend:dev --target dev .

Produção (mínima e segura):
  docker build -t chatear-backend:prod --target prod .


## Antigo Dockerfile
# ===============================
# Stage 1: Build
# ===============================
FROM golang:1.25.3-alpine AS builder

# Instala dependências necessárias
RUN apk add --no-cache git ca-certificates build-base

# Define diretório de trabalho
WORKDIR /app

# Copia e baixa dependências — isolado para melhor cache
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copia o restante do código
COPY . .

# Compila todos os binários de uma vez (mais rápido)
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/chatear-api ./cmd/api && \
    CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/notification_worker ./cmd/worker/notification_worker.go && \
    CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/user_delete_worker ./cmd/worker/user_delete_worker.go && \
    CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/user_hard_delete_worker ./cmd/worker/user_hard_delete_worker.go && \
    CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/user_permanent_deletion_scheduler_worker ./cmd/worker/user_permanent_deletion_scheduler_worker.go && \
    CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/user_registered_worker ./cmd/worker/user_registered_worker.go && \
    CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/password_recovery_worker ./cmd/worker/password_recovery_worker.go

# ===============================
# Stage 2: Production
# ===============================
FROM alpine:3.20

# Instala certificados SSL
RUN apk add --no-cache ca-certificates tzdata

# Cria diretório do app
WORKDIR /app

# Copia binários do estágio anterior
COPY --from=builder /app/bin/* ./
COPY --from=builder /app/internal/notification/infrastructure/templates ./internal/notification/infrastructure/templates

# Expõe porta da API
EXPOSE 8080

# Variável padrão de execução
ENV APP_BIN=chatear-api

# Permite substituir APP_BIN no docker-compose
ENTRYPOINT ["sh", "-c", "./$APP_BIN"]




## Antigo dokcer-compose.events.yml
x-common-env: &common-env
  env_file:
    - .env
  volumes:
    - ./.env:/app/.env:ro
  depends_on:
    nats:
      condition: service_healthy
    redis:
      condition: service_healthy
  networks:
    - my_ipv6_network
  build:
    context: .
    dockerfile: Dockerfile

services:
  api:
    <<: *common-env
    container_name: chatear-backend-api
    ports:
      - "8080:8080"
    environment:
      - APP_BIN=chatear-api
    command: ["sh", "-c", "./chatear-api"]

  notification-worker:
    <<: *common-env
    container_name: chatear-notification-worker
    environment:
      - APP_BIN=notification_worker
    command: ["sh", "-c", "./notification_worker"]

  user-delete-worker:
    <<: *common-env
    container_name: chatear-user-delete-worker
    environment:
      - APP_BIN=user_delete_worker
    command: ["sh", "-c", "./user_delete_worker"]

  user-hard-delete-worker:
    <<: *common-env
    container_name: chatear-user-hard-delete-worker
    environment:
      - APP_BIN=user_hard_delete_worker
    command: ["sh", "-c", "./user_hard_delete_worker"]

  user-permanent-deletion-scheduler-worker:
    <<: *common-env
    container_name: chatear-user-permanent-deletion-scheduler-worker
    environment:
      - APP_BIN=user_permanent_deletion_scheduler_worker
    command: ["sh", "-c", "./user_permanent_deletion_scheduler_worker"]

  user-registered-worker:
    <<: *common-env
    container_name: chatear-user-registered-worker
    environment:
      - APP_BIN=user_registered_worker
    command: ["sh", "-c", "./user_registered_worker"]

  password-recovery-worker:
    <<: *common-env
    container_name: chatear-password-recovery-worker
    environment:
      - APP_BIN=password_recovery_worker
    command: ["sh", "-c", "./password_recovery_worker"]

  nats:
    image: nats:2.10-alpine
    container_name: chatear-backend-nats
    ports:
      - "4222:4222"
      - "8222:8222"
    command:
      - "--jetstream"
      - "--http_port=8222"
      - "--port=4222"
    volumes:
      - nats-data:/data
    networks:
      - my_ipv6_network
    healthcheck:
      test: ["CMD", "wget", "--spider", "-q", "http://localhost:8222/healthz"]
      interval: 5s
      timeout: 5s
      retries: 10
      start_period: 5s

  redis:
    image: redis:7-alpine
    container_name: chatear-backend-redis
    ports:
      - "6379:6379"
    volumes:
      - redis-data:/data
    networks:
      - my_ipv6_network
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 5s
      retries: 10
      timeout: 5s
      start_period: 5s

volumes:
  nats-data:
  redis-data:

networks:
  my_ipv6_network:
    external: true

