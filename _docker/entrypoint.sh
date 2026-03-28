#!/bin/sh
set -e

MAX_RETRIES=10
RETRY_COUNT=0

# Путь к сертификатам
if [ "$APP_ENV" = "prod" ]; then
  CERTS_DIR="/certs-cloud"
else
  CERTS_DIR="/certs"
fi

echo "Using environment: $APP_ENV"
echo "HOTELS_SERVICE_PORT: $HOTELS_SERVICE_PORT"

# Ожидание доступности порта
echo "⏳ Waiting for CockroachDB at ${HOTELS_COCKROACH_HOST}:${HOTELS_COCKROACH_PORT_INNER}..."
until nc -z "$HOTELS_COCKROACH_HOST" "$HOTELS_COCKROACH_PORT_INNER"; do
  RETRY_COUNT=$((RETRY_COUNT + 1))

  echo "✅ Attempt $RETRY_COUNT"
  if [ "$RETRY_COUNT" -ge "$MAX_RETRIES" ]; then
    echo "❌ Failed to connect to CockroachDB after ${MAX_RETRIES} attempts. Exiting."
    exit 1
  fi
  sleep 1
done
echo "✅ CockroachDB is available!"

# Проверяем и создаём пользователя и базу через root
echo "🛠 Ensuring user '${HOTELS_COCKROACH_USER}' and database '${HOTELS_COCKROACH_DB_NAME}' exist..."
cockroach sql \
  --certs-dir="$CERTS_DIR" \
  --host="$HOTELS_COCKROACH_HOST" \
  --port="$HOTELS_COCKROACH_PORT_INNER" \
  --user=root \
  --execute="
    CREATE USER IF NOT EXISTS ${HOTELS_COCKROACH_USER};
    CREATE DATABASE IF NOT EXISTS ${HOTELS_COCKROACH_DB_NAME};
    GRANT ALL ON DATABASE ${HOTELS_COCKROACH_DB_NAME} TO ${HOTELS_COCKROACH_USER};
  "
echo "✅ User and database ready."

# Проверка соединения уже от HOTELS_COCKROACH_USER после создания пользователя
echo "🔐 Verifying connection to CockroachDB... as '${HOTELS_COCKROACH_USER}'..."
cockroach sql \
  --certs-dir="$CERTS_DIR" \
  --host="$HOTELS_COCKROACH_HOST" \
  --port="$HOTELS_COCKROACH_PORT_INNER" \
  --user="$HOTELS_COCKROACH_USER" \
  --database="$HOTELS_COCKROACH_DB_NAME" \
  --execute="SELECT 1;"

if [ $? -ne 0 ]; then
  echo "❌ Unable to connect to CockroachDB."
  exit 1
fi

# Путь к корню микросервиса hotels-service
HOTELS_SERVICE_ROOT="$(cd "$(dirname "$0")/.." && pwd)"

echo "📦 Applying database migrations..."
# Передаем переменные напрямую через окружение и ставим флаг
USE_ENTRYPOINT_ENV=1 \
HOTELS_COCKROACH_HOST="$HOTELS_COCKROACH_HOST" \
HOTELS_COCKROACH_PORT_INNER="$HOTELS_COCKROACH_PORT_INNER" \
HOTELS_COCKROACH_USER="$HOTELS_COCKROACH_USER" \
HOTELS_COCKROACH_PASSWORD="$HOTELS_COCKROACH_PASSWORD" \
HOTELS_COCKROACH_DB_NAME="$HOTELS_COCKROACH_DB_NAME" \
CERTS_DIR="$CERTS_DIR" \
"$HOTELS_SERVICE_ROOT/db/migrate.sh"

# Database seeding
go run "$HOTELS_SERVICE_ROOT/cmd/seed/main.go"

export AIR_LOG_LEVEL=debug
# Запуск основного приложения
echo "🚀 Starting hotels-service..."
exec air -c .air.toml