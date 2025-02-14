# Используем официальный образ Golang
FROM golang:1.23.6-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем модули и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Запускаем тесты перед сборкой
RUN go test ./...

# Собираем бинарник
RUN go build -o hotel-service main.go

# Финальный контейнер
FROM alpine:latest  

WORKDIR /root/
COPY --from=builder /app/hotel-service .

# Expose порта
EXPOSE 9064

# Запускаем сервис
CMD ["./hotel-service"]