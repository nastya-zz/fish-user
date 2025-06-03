LOCAL_BIN:=$(CURDIR)/bin

functional-tests:
	go test ./...

generate:
	make generate-migration

generate-migration:
	./bin/goose create ____ sql

up-test-env:
	docker-compose -f docker-compose.test.yaml --env-file .env.test up -d --build

