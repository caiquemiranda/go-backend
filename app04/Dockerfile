FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copiar apenas o arquivo go.mod primeiro para aproveitar o cache do Docker
COPY go.mod ./
RUN go mod download

# Agora copiar o restante dos fontes
COPY . .

# Compilar o aplicativo
RUN CGO_ENABLED=0 GOOS=linux go build -o app04 .

# Imagem final, usando alpine mínimo
FROM alpine:3.17

WORKDIR /app

# Instalar certificados e fuso horário
RUN apk --no-cache add ca-certificates tzdata

# Copiar o binário compilado
COPY --from=builder /app/app04 .

# Expor porta
EXPOSE 8080

# Comando para executar o aplicativo
CMD ["./app04"] 