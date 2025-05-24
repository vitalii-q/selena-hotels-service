# Stage 1: Build the Go binary
FROM golang:1.23.6-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

#RUN go test ./...

RUN go build -o hotel-service main.go

# Stage 2: Runtime image based on Debian slim (glibc included)
FROM debian:bullseye-slim

WORKDIR /root/

# Устанавливаем curl и netcat (nc), необходимые для скрипта ожидания и загрузки cockroach cli
RUN apt-get update && apt-get install -y curl netcat ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Установка cockroach CLI
RUN curl -s https://binaries.cockroachdb.com/cockroach-v22.2.7.linux-amd64.tgz | tar -xz \
    && cp -i cockroach-v22.2.7.linux-amd64/cockroach /usr/local/bin/ \
    && chmod +x /usr/local/bin/cockroach \
    && rm -rf cockroach-v22.2.7.linux-amd64*

# Копируем бинарник приложения
COPY --from=builder /app/hotel-service .

# Копируем скрипт ожидания базы и даём права на исполнение
COPY _docker/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

EXPOSE 9064

ENTRYPOINT ["./entrypoint.sh"]
