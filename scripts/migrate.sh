#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ENV_FILE="$SCRIPT_DIR/../.env"

if [ -f "$ENV_FILE" ]; then
  source "$ENV_FILE"
else
  echo "❌ .env file not found at $ENV_FILE"
  exit 1
fi

if [ -z "$DATABASE_URI" ]; then
  echo "❌ DATABASE_URI not set in .env"
  exit 1
fi

DB_URL="$DATABASE_URI"
MIGRATIONS_PATH="$SCRIPT_DIR/../migrations"

case "$1" in

  up)
    echo "Running migrations UP..."
    migrate -path "$MIGRATIONS_PATH" -database "$DB_URL" up
    ;;

  down)
    echo "Running migrations DOWN..."
    migrate -path "$MIGRATIONS_PATH" -database "$DB_URL" down
    ;;

  create)
    if [ -z "$2" ]; then
      echo "❌ Please provide a migration name"
      echo "Usage: ./migrate.sh create migration_name"
      exit 1
    fi
    echo "Creating migration: $2"
    migrate create -ext sql -dir "$MIGRATIONS_PATH" -seq "$2"
    ;;

  force)
    if [ -z "$2" ]; then
      echo "❌ Please provide version number"
      echo "Usage: ./migrate.sh force version_number"
      exit 1
    fi
    echo "Forcing migration version to $2"
    migrate -path "$MIGRATIONS_PATH" -database "$DB_URL" force "$2"
    ;;

  reset)
    echo "⚠️  This will wipe all data and re-run all migrations."
    read -p "Continue? [y/N] " confirm
    [[ "$confirm" == "y" ]] || exit 1
    echo "Resetting database..."
    migrate -path "$MIGRATIONS_PATH" -database "$DB_URL" down -all && \
    migrate -path "$MIGRATIONS_PATH" -database "$DB_URL" up
    ;;

  *)
    echo "Usage:"
    echo "  ./migrate.sh up"
    echo "  ./migrate.sh down"
    echo "  ./migrate.sh create <name>"
    echo "  ./migrate.sh force <version>"
    echo "  ./migrate.sh reset"
    ;;

esac
