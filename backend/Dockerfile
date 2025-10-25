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

