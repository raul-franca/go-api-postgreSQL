
createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path sql/migrations -database "postgresql://postgres:pass@localhost:5432/todos?sslmode=disable" -verbose up

migratedown:
	migrate -path=sql/migrations -database "postgresql://postgres:pass@localhost:5432/todos?sslmode=disable" -verbose down

.PHONY: migrate migratedown createmigration

# -database "postgresql://postgres:pass@localhost:5432/postgres" -path sql/migrations up