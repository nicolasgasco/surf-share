set -euo pipefail

cd /go/app

DB_HOST="${DB_HOST:-db}"
DB_PORT="${DB_PORT:-5432}"
until pg_isready -h "$DB_HOST" -p "$DB_PORT" >/dev/null 2>&1; do
  echo "waiting for database at $DB_HOST:$DB_PORT..."
  sleep 1
done

if command -v air >/dev/null 2>&1; then
  exec air -c .air.toml
else
  exec go run main.go
fi