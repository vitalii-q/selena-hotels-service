#!/bin/bash

set -e

DB_HOST="localhost"
DB_PORT="9264"
DB_USER="hotels_user"
DB_NAME="hotels_db"
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
