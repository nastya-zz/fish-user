#!/bin/bash
source .env.test

export MIGRATION_DSN="host=pg_test port=5432 dbname=$PG_DATABASE_NAME user=$PG_USER password=$PG_PASSWORD sslmode=disable"

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v