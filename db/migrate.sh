#!/bin/bash

# Launch DB migration:
# docker exec -it hotels-service bash /app/db/migrate.sh

set -e # Падение скрипта при любой ошибке

# Если переменная USE_ENTRYPOINT_ENV установлена, запуск миграций идет из /_docker/entrypoint.sh
if [ -z "$USE_ENTRYPOINT_ENV" ]; then
    echo "🔹 Loading .env file..."
    set -o allexport
    source ".env"
    set +o allexport
else
    echo "🔹 Using environment variables passed from entrypoint.sh"
fi

# Определим, где мы запускаемся: в контейнере или на хосте
if grep -q docker /proc/1/cgroup || [ -f /.dockerenv ]; then
  echo "🧱 Running inside Docker container"
  DB_HOST=${HOTELS_COCKROACH_HOST}
  DB_PORT=${HOTELS_COCKROACH_PORT_INNER}
else
  echo "💻 Running on host machine"
  DB_HOST=${LOCALHOST}
  DB_PORT=${HOTELS_COCKROACH_PORT}
fi

DB_USER="${HOTELS_COCKROACH_USER}"
DB_PASSWORD="${HOTELS_COCKROACH_PASSWORD}"
DB_NAME="${HOTELS_COCKROACH_DB_NAME}"
MIGRATIONS_DIR="db/migrations"
CERTS_DIR="${CERTS_DIR:-/certs}"

echo "DB_HOST=$DB_HOST"
echo "DB_PORT=$DB_PORT"
echo "CERTS_DIR=$CERTS_DIR"

echo "Applying migrations from $MIGRATIONS_DIR..."

shopt -s nullglob
FILES=("$MIGRATIONS_DIR"/*.up.sql)
shopt -u nullglob

if [ ${#FILES[@]} -eq 0 ]; then
    echo "No migration files found."
    exit 1
fi

DB_URL="postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=verify-full"
for file in "${FILES[@]}"; do
    echo "Applying migration file: $file"
    
    if ! cockroach sql \
        --certs-dir="$CERTS_DIR" \
        --url="$DB_URL" \
        --file="$file"; then
        echo "❌ Error applying migration: $file"
        exit 1
    fi
done

echo "Migrations applied successfully!"
