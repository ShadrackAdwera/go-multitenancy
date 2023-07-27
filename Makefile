DB_URL=postgres://postgres:password@localhost:5431/go_multitenancy?sslmode=disable
TEST_DB_URL=postgres://postgres:password@localhost:5431/test_go_multitenancy?sslmode=disable

migrate_create:
	migrate create -ext sql -dir db/migrations -seq ${MIGRATE_NAME}
migrate_up:
	migrate -path db/migrations -database "${TEST_DB_URL}" -verbose up
migrate_down:
	migrate -path db/migrations -database "${TEST_DB_URL}" -verbose down
sqlc:
	sqlc generate --file internal/db/sqlc.yaml
tests:
	go test -v -cover ./...
mocks:
	mockgen -package mockdb --destination pkg/mocks/store.go github.com/ShadrackAdwera/go-multitenancy/internal/sqlc TxStore
start:
	go run main.go

.PHONY: migrate_create migrate_up migrate_down sqlc tests mocks start