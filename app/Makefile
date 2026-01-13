.PHONY: test-integration

test-integration:
	@export DB_PASSWORD=$$(cat /run/secrets/db_password) && \
	go test ./test/integration -v
