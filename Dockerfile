# Stage 1: Build the Go binary
FROM golang:1.23.6 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .

#RUN go test ./...

RUN go build -o hotels-service main.go

# Stage 2: Runtime 
FROM golang:1.23.6

WORKDIR /app

# Устанавливаем curl и netcat (nc), необходимые для скрипта ожидания и загрузки cockroach cli
RUN apt-get update && apt-get install -y curl netcat-openbsd ca-certificates postgresql-client \
    && rm -rf /var/lib/apt/lists/*

# Установка cockroach CLI
RUN curl -s https://binaries.cockroachdb.com/cockroach-v22.2.7.linux-amd64.tgz | tar -xz \
    && cp -i cockroach-v22.2.7.linux-amd64/cockroach /usr/local/bin/ \
    && chmod +x /usr/local/bin/cockroach \
    && rm -rf cockroach-v22.2.7.linux-amd64*

# Копируем скомпилированное приложение и air
COPY --from=builder /app .
#COPY --from=builder /app/hotels-service /usr/local/bin/hotels-service # 
COPY --from=builder /go/bin/air /usr/local/bin/air

EXPOSE ${HOTELS_SERVICE_PORT}

ENTRYPOINT ["./_docker/entrypoint.sh"]
