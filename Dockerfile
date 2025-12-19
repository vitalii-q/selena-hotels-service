# hotels-service/Dockerfile

# --- Start microservice
# docker build --no-cache --platform=linux/amd64 -t selena-hotels-service:latest .
#
# docker run -d --name hotels-service --env-file .env -p 9064:9064 --network selena-dev_app_network -v $(pwd):/app selena-hotels-service:latest
# -v $(pwd):/app/hotels-service — mount the local sources into the container

# --- Start DB for microservice
# 

# Stage 1: Build the Go binary
FROM golang:1.25 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/air-verse/air@v1.63.0

COPY . .

#RUN go test ./...

RUN go build -o hotels-service main.go

# Stage 2: Runtime 
FROM golang:1.25

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
COPY --from=builder /go/bin/air /usr/local/bin/air

COPY db/certs /certs

EXPOSE ${HOTELS_SERVICE_PORT}

ENTRYPOINT ["./_docker/entrypoint.sh"]
