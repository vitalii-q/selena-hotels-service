#!/bin/sh
set -e

MAX_RETRIES=10
RETRY_COUNT=0

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

# Verifying SQL connection
echo "🔐 Verifying connection to CockroachDB..."
cockroach sql --insecure --host="$HOTELS_COCKROACH_HOST" --port="$HOTELS_COCKROACH_PORT_INNER" --execute="SELECT 1;"

if [ $? -ne 0 ]; then
  echo "❌ Unable to connect to CockroachDB."
  exit 1
fi

# Проверка и создание пользователя и базы

echo "🔍 Checking if database '${HOTELS_COCKROACH_DB_NAME}' exists..."

# Обрати внимание: правильное имя колонки — database_name (нижнее подчеркивание)
if ! cockroach sql --insecure --host="$HOTELS_COCKROACH_HOST" --port="$HOTELS_COCKROACH_PORT_INNER" \
    --execute="SELECT 1 FROM [SHOW DATABASES] WHERE database_name = '${HOTELS_COCKROACH_DB_NAME}';" | grep -q "1"; then

  echo "🛠 Creating user '${HOTELS_COCKROACH_USER}' and database '${HOTELS_COCKROACH_DB_NAME}'..."

  # В insecure режиме пароль не задаём, поэтому просто создаём пользователя без пароля
  cockroach sql --insecure --host="$HOTELS_COCKROACH_HOST" --port="$HOTELS_COCKROACH_PORT_INNER" --execute="
    CREATE USER IF NOT EXISTS ${HOTELS_COCKROACH_USER};
    CREATE DATABASE IF NOT EXISTS ${HOTELS_COCKROACH_DB_NAME};
    GRANT ALL ON DATABASE ${HOTELS_COCKROACH_DB_NAME} TO ${HOTELS_COCKROACH_USER};
  "

  echo "✅ User and database created."
else
  echo "📦 Database '${HOTELS_COCKROACH_DB_NAME}' already exists."
fi

# Запуск основного приложения
echo "🚀 Starting hotel-service..."
exec ./hotel-service
