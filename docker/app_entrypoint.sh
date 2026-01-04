set -euo pipefail

# Required pgx
export DB_PASSWORD=$(cat /run/secrets/db_password)

cd /go/app

echo "Installing pgx and dependencies..."
go get github.com/jackc/pgx/v5
go get github.com/jackc/pgx/v5/pgconn
go get github.com/jackc/pgx/v5/pgtype

echo "Tidying go modules..."
go mod tidy

echo "Downloading all dependencies..."
go mod download

echo "Starting the application..."
if command -v air >/dev/null 2>&1; then
  exec air -c .air.toml
else
  exec go run main.go
fi