#!/bin/bash
set -e

if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

if [ -z "$DATABASE_URL" ]; then
    echo "Error: DATABASE_URL is not set in .env"
    exit 1
fi

MIGRATIONS_PATH="migrations/"

echo "Applying migrations from $MIGRATIONS_PATH..."
migrate -path "$MIGRATIONS_PATH" -database "$DATABASE_URL" up

echo "Migrations applied successfully."
