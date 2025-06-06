LOCAL_BIN:=$(CURDIR)/bin


up-test-app-with_test-env:
	make up-test-env
	make run-app-test-env

functional-tests:
	go test ./tests/...

generate:
	make generate-migration

generate-migration:
	./bin/goose create ____ sql

up-test-env:
	docker-compose -f docker-compose.test.yaml --env-file .env.test up -d --build

run-app-test-env:
	go run cmd/main.go --env=test