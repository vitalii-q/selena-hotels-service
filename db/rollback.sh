#!/bin/bash
set -euo pipefail

echo "🧭 Current working directory: $(pwd)"
echo "📁 Script is located in: $(cd "$(dirname "$0")" && pwd)"

APP_ROOT="/app"
ENV_FILE="$APP_ROOT/.env"

if [ -f "$ENV_FILE" ]; then
  echo "🔑 Loading env from $ENV_FILE"
  set -o allexport
  source "$ENV_FILE"
  set +o allexport
fi

: "${HOTELS_COCKROACH_HOST:?}"
: "${HOTELS_COCKROACH_PORT_INNER:?}"
: "${HOTELS_COCKROACH_USER:?}"
: "${HOTELS_COCKROACH_PASSWORD:?}"
: "${HOTELS_COCKROACH_DB_NAME:?}"
: "${DB_SSLMODE:?}"

DB_HOST="$HOTELS_COCKROACH_HOST"
DB_PORT="$HOTELS_COCKROACH_PORT_INNER"
DB_USER="$HOTELS_COCKROACH_USER"
DB_PASSWORD="$HOTELS_COCKROACH_PASSWORD"
DB_NAME="$HOTELS_COCKROACH_DB_NAME"
SSL_MODE="$DB_SSLMODE"

CERTS_DIR="/certs"
SSL_ROOT_CERT="$CERTS_DIR/ca.crt"
SSL_CERT="$CERTS_DIR/client.${DB_USER}.crt"
SSL_KEY="$CERTS_DIR/client.${DB_USER}.key"

MIGRATIONS_DIR="$APP_ROOT/db/migrations"

echo "Rolling back migrations from $MIGRATIONS_DIR..."
echo "→ DB Host: $DB_HOST:$DB_PORT"
echo "→ SSL Mode: $SSL_MODE, certs: $CERTS_DIR"

shopt -s nullglob
FILES=("$MIGRATIONS_DIR"/*.down.sql)
if [ ${#FILES[@]} -eq 0 ]; then
  echo "No .down.sql migration files found in $MIGRATIONS_DIR"
  exit 0
fi

# Is important for CockroachDB
export PGCLIENTENCODING=UTF8

for file in $(ls -r "$MIGRATIONS_DIR"/*.down.sql); do
  echo "Reverting migration: $file"
  psql "host=$DB_HOST \
        port=$DB_PORT \
        dbname=$DB_NAME \
        user=$DB_USER \
        password=$DB_PASSWORD \
        sslmode=$SSL_MODE \
        sslrootcert=$SSL_ROOT_CERT \
        sslcert=$SSL_CERT \
        sslkey=$SSL_KEY" \
        -f "$file"
done

echo "Rollback completed successfully!"
