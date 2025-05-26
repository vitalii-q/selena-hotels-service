#!/bin/bash
echo "🧭 Current working directory: $(pwd)"
echo "📁 Script is located in: $(cd "$(dirname "$0")" && pwd)"

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

echo "Rolling back migrations from $MIGRATIONS_DIR..."

FILES=("$MIGRATIONS_DIR"/*.down.sql)

# Проверка, есть ли .down.sql файлы
if [ ! -e "${FILES[0]}" ]; then
    echo "No .down.sql migration files found in $MIGRATIONS_DIR"
    exit 0
fi

# Применение миграций в обратном порядке
for file in $(ls -r "$MIGRATIONS_DIR"/*.down.sql); do
    echo "Reverting migration: $file"
    if ! PGPASSWORD="postgres" psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -f "$file"; then
        echo "Error reverting migration: $file"
        exit 1
    fi
done

echo "Rollback completed successfully!"
