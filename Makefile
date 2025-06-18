# ===== Include environment variables =====
ifneq (,$(wildcard .env))
	include .env
	export
endif

# ===== Default configs =====
MIGRATE=migrate
AUTH_MIGRATIONS=services/auth-service/migrations
AUTH_DB_URL=$(POSTGRES_URL)

# ===== AUTH SERVICE MIGRATIONS =====

auth-migrate-up:
	$(MIGRATE) -path $(AUTH_MIGRATIONS) -database "$(AUTH_DB_URL)" up

auth-migrate-down:
	$(MIGRATE) -path $(AUTH_MIGRATIONS) -database "$(AUTH_DB_URL)" down 1

auth-migrate-new:
	@if [ -z "$(name)" ]; then \
		echo "Please provide migration name: make auth-migrate-new name=create_xyz"; \
		exit 1; \
	fi
	$(MIGRATE) create -ext sql -dir $(AUTH_MIGRATIONS) -seq $(name)

# ===== ORDER SERVICE MIGRATIONS =====

ORDER_MIGRATIONS=services/order-service/migrations
ORDER_DB_URL=$(POSTGRES_URL)  # หรือใช้ตัวแปรแยก เช่น ORDER_POSTGRES_URL

order-migrate-up:
	$(MIGRATE) -path $(ORDER_MIGRATIONS) -database "$(ORDER_DB_URL)" up

order-migrate-down:
	$(MIGRATE) -path $(ORDER_MIGRATIONS) -database "$(ORDER_DB_URL)" down 1

order-migrate-new:
	@if [ -z "$(name)" ]; then \
		echo "Please provide migration name: make order-migrate-new name=create_xyz"; \
		exit 1; \
	fi
	$(MIGRATE) create -ext sql -dir $(ORDER_MIGRATIONS) -seq $(name)


# ===== MongoDB schema validation =====

product-setup-validator:
	node services/product-service/scripts/setupValidator.js

# ===== SEED (optional) =====
# product-seed:
# 	node services/product-service/seed.js

# ===== TEST ALL (optional) =====
# test-all:
# 	go test ./...
# 	npm test --prefix services/product-service
