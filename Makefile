.PHONY: test migrate

ci: test lint

run:
	go run main.go

generate:
	/bin/sh bin/generate.sh

migrate-up:
	go run main.go migrate up

migrate-down:
	go run main.go migrate down

swag:
	swag init -g internal/app/app.go --pd --overridesFile override.swag

local-compose:
	docker compose -f build/local/docker-compose.local.yml up --detach --build

mock:
	mockery

test:
	go test ./... -p 1 -cover -coverprofile=coverage.out -coverpkg=./...

coverage:
	touch coverage.out
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

lint:
	golangci-lint run

# use with caution
# can delete nolint comments
# can mess with imports
lintfix:
	golangci-lint run --fix

fmt:
	gofumpt -l -w .
	gci write . --skip-generated -s standard -s default -s "prefix(cf-pcred-terminatepledge-backend)" -s blank -s dot

migrate:
	migrate -path migrations/agreement -database postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@$$POSTGRES_HOST:$$POSTGRES_PORT/$$POSTGRES_DB?sslmode=disable up

new_migration_help:
	echo "migrate create -ext sql -dir migrations/data -seq {name}"