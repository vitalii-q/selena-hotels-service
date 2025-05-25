#!/bin/bash

DB_HOST="localhost"
DB_PORT="9264"
DB_USER="hotels_user"
DB_NAME="hotels_db"
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
