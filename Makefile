LOCAL_BIN:=$(CURDIR)/bin

functional-tests:
	go test ./...

generate:
	make generate-migration

generate-migration:
	./bin/goose create ____ sql