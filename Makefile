# Makefile (for CI)
# ------------------
# Place in project root as `Makefile`

# Build production images
dev-up:
	docker-compose -f docker-compose.dev.yml up --build

dev-down:
	docker-compose -f docker-compose.dev.yml down

# データベースマイグレーションを実行
migrate:
	docker cp backend/migrations/20250412093819_init.up.sql supabase-db:/tmp/init.sql
	docker exec -i supabase-db psql -U postgres -d postgres -c "\i /tmp/init.sql"

prod-up:
	docker-compose -f docker-compose.prod.yml up -d --build

prod-down:
	docker-compose -f docker-compose.prod.yml down