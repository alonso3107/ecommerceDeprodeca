.PHONY: dev build run db-up db-down db-reset migrate-up migrate-down sqlc-gen tidy

# ─── Desarrollo ─────────────────────────────────────────────
dev:
	go run ./cmd/api

build:
	go build -o bin/api ./cmd/api

run: build
	./bin/api

# ─── Base de Datos (Docker) ─────────────────────────────────
db-up:
	docker-compose up -d postgres redis

db-down:
	docker-compose down

db-reset: db-down
	docker-compose down -v
	docker-compose up -d postgres redis

# ─── Migraciones (Goose) ────────────────────────────────────
migrate-up:
	goose -dir db/migrations postgres "$$DATABASE_URL" up

migrate-down:
	goose -dir db/migrations postgres "$$DATABASE_URL" down

migrate-status:
	goose -dir db/migrations postgres "$$DATABASE_URL" status

# ─── SQLC ───────────────────────────────────────────────────
sqlc-gen:
	sqlc generate -f db/sqlc.yaml

# ─── Utilidades ─────────────────────────────────────────────
tidy:
	go mod tidy
