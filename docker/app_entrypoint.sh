set -euo pipefail

# Required pgx
DB_PASSWORD=$(cat /run/secrets/db_password)
export DB_PASSWORD

cd /go/app

echo "Installing make command..."
apk add --no-cache make

echo "Installing pgx and dependencies..."
go get github.com/jackc/pgx/v5
go get github.com/jackc/pgx/v5/pgconn
go get github.com/jackc/pgx/v5/pgtype
go get github.com/jackc/pgx/v5/pgxpool
go get github.com/georgysavva/scany/v2

echo "Installing air for hot reload..."
go install github.com/air-verse/air@latest

echo "Tidying go modules..."
go mod tidy

echo "Downloading all dependencies..."
go mod download

echo "Starting the application..."
if command -v air >/dev/null 2>&1; then
  echo "Starting in development mode with air..."
  exec air -c .air.toml
else
  echo "Starting in production mode..."
  exec go run main.go
fi