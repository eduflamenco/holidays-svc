# ---- Build Stage ----
FROM golang:1.24-alpine AS builder

# Configuración del entorno
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copiar código y compilar
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/holidays-svc \
                                                           ./cmd/holidays/main.go

# ---- Runtime Stage ----
FROM alpine:3.18

# Copiar binario desde el builder
WORKDIR /app
COPY --from=builder /app/bin/holidays-svc /app/holidays-svc
COPY --from=builder /app/app.env .

RUN chmod +x /app/holidays-svc

# Variables de entorno (ajustables)
ENV PORT=8080
EXPOSE $PORT

CMD ["/app/holidays-svc"]