# Dockerfile
FROM golang:1.23-alpine AS builder

# Устанавливаем зависимости для сборки
RUN apk add --no-cache git ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы модулей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./backend/cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./backend/cmd/main.go
# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из builder stage
COPY --from=builder /app/main .
# Копируем конфигурационные файлы
COPY ./backend/cmd/conn_config.env .
COPY ./backend/cmd/db_config.env .
COPY ./backend/cmd/.env .

# Экспонируем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]