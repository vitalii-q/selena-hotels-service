#!/bin/bash

set -e # Падение скрипта при любой ошибке

# Подключаем переменные окружения из .env
set -o allexport
source ".env"
set +o allexport
#cat ".env"

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
DB_NAME="${HOTELS_COCKROACH_DB_NAME}"
MIGRATIONS_DIR="db/migrations"

echo "Applying migrations from $MIGRATIONS_DIR..."

shopt -s nullglob
FILES=("$MIGRATIONS_DIR"/*.up.sql)
shopt -u nullglob

if [ ${#FILES[@]} -eq 0 ]; then
    echo "No migration files found."
    exit 1
fi

for file in "${FILES[@]}"; do
    echo "Applying migration: $file"
    if ! PGPASSWORD="postgres" psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -f "$file"; then
        echo "Error applying migration: $file"
        exit 1
    fi
done

echo "Migrations applied successfully!"
