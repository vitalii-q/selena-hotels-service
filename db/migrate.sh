#!/bin/bash

set -e # Падение скрипта при любой ошибке

# Абсолютный путь до корня проекта selena-dev (два уровня выше скрипта)
ROOT_DIR="$(cd "$(dirname "$0")/../../" && pwd)"

# Подключаем переменные окружения из .env
set -o allexport
source "$ROOT_DIR/.env"
set +o allexport

DB_HOST="${LOCALHOST}"
DB_PORT="${HOTELS_COCKROACH_PORT}"
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
