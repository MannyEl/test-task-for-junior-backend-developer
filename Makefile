.PHONY: migrate-up
migrate-up:
	migrate -path ./migrations -database ${DATABASE_DSN} up $(filter-out $@, ${MAKECMDGOALS})

