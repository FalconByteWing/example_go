# Database configuration
DB_URL?=postgres://postgres:postgres@localhost:5432/example_db?sslmode=disable

# Directory creation command based on OS
ifeq ($(OS),Windows_NT)
	MKDIR = cmd /c if not exist
	RUN_WITH_ENV = powershell -Command "Get-Content .env | ForEach-Object { $$line = $$_ -split '=',2; if ($$line.Length -eq 2) { [System.Environment]::SetEnvironmentVariable($$line[0], $$line[1]) }}; 
else
	MKDIR = mkdir -p
	RUN_WITH_ENV = set -a && source .env && set +a &&
endif


.PHONY: swagger-tsp-service-payout
swagger-tsp-service-payout:
	$(MKDIR) "services\tsp-service-payout\generated" mkdir "services\tsp-service-payout\generated"
	swagger generate server -A tsp-payout-service -f services/tsp-service-payout/api/swagger.yaml -t services/tsp-service-payout/generated


.PHONY: swagger
swagger: swagger-tsp-service-payout

# Build commands
.PHONY: build
build:
	go build -o bin/tsp-payout-service services/tsp-service-payout/cmd/main.go

.PHONY: run-tsp-service-payout
run-tsp-service-payout:
ifeq ($(OS),Windows_NT)
	$(RUN_WITH_ENV) go run services/tsp-service-payout/cmd/main.go"
else
	./scripts/run-tsp-service-payout.sh
endif

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: create-db
create-db:
ifeq ($(OS),Windows_NT)
	@powershell -Command "$$env:PGPASSWORD='postgres'; psql -h localhost -U postgres -c \"CREATE DATABASE example_db;\""
else
	PGPASSWORD=postgres psql -h localhost -U postgres -c "CREATE DATABASE example_db;"
endif

.PHONY: drop-db
drop-db:
ifeq ($(OS),Windows_NT)
	@powershell -Command "$$env:PGPASSWORD='postgres'; psql -h localhost -U postgres -c \"DROP DATABASE IF EXISTS example_db;\""
else
	PGPASSWORD=postgres psql -h localhost -U postgres -c "DROP DATABASE IF EXISTS example_db;"
endif

# Migration commands
.PHONY: migrate-create
migrate-create:
	migrate create -ext sql -dir services/tsp-service-payout/internal/db/migrations -seq $(name)

.PHONY: migrate-up
migrate-up:
	migrate -path services/tsp-service-payout/internal/db/migrations -database "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	migrate -path services/tsp-service-payout/internal/db/migrations -database "$(DB_URL)" down

.PHONY: migrate-force
migrate-force:
	migrate -path services/tsp-service-payout/internal/db/migrations -database "$(DB_URL)" force $(version)

# Update init-db target to include migrations
.PHONY: init-db
init-db: create-db migrate-up sqlc

.PHONY: swagger-serve-tsp-service-payout
swagger-serve-tsp-service-payout:
	swagger serve -F swagger services/tsp-service-payout/api/swagger.yaml 

# Add this at the top of your Makefile, after the DB_URL definition
ifeq ($(OS),Windows_NT)
	PATHSEP=\\
else
	PATHSEP=/
endif 