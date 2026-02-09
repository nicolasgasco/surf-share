set -euo pipefail

cd /go/app

echo "Installing make command..."
apk add --no-cache make

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
