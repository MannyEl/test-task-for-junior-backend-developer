.PHONY: migrate-up
migrate-up:
	migrate -path ./migrations -database ${DATABASE_DSN} up $(filter-out $@, ${MAKECMDGOALS})

.PHONY: seed
seed:
	go run ./cmd/seed/seed.go