#!/bin/bash

# Conduit Database Utility Script
# Provides common database operations for development

# Load environment variables from .env file
if [ -f .env ]; then
    export $(cat .env | grep -v '^#' | xargs)
fi

DB_USER=${CONDUIT_DB_USER:-conduituser}
DB_PASSWORD=${CONDUIT_DB_PASSWORD:-conduitpass}
DB_NAME=${CONDUIT_DB_NAME:-conduitdb}
DB_HOST=${CONDUIT_DB_HOST:-localhost}

PGCMD="PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME"

case "$1" in
    "connect")
        echo "🔗 Connecting to database..."
        PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME
        ;;
    "tables")
        echo "📋 Listing all tables:"
        PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "\dt"
        ;;
    "users")
        echo "👥 Users table structure:"
        PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "\d users"
        echo ""
        echo "👥 Current users:"
        PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "SELECT id, username, email, created_at FROM users ORDER BY created_at DESC LIMIT 10;"
        ;;
    "reset")
        echo "⚠️  WARNING: This will drop and recreate the database!"
        read -p "Are you sure? (y/N): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            echo "🗑️  Dropping database..."
            sudo -u postgres psql -c "DROP DATABASE IF EXISTS $DB_NAME;"
            echo "🆕 Creating fresh database..."
            sudo -u postgres psql -c "CREATE DATABASE $DB_NAME OWNER $DB_USER;"
            sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;"
            echo "✅ Database reset complete. Run your app to recreate tables."
        else
            echo "❌ Operation cancelled."
        fi
        ;;
    "status")
        echo "📊 Database Status:"
        echo "  Host: $DB_HOST"
        echo "  Database: $DB_NAME"
        echo "  User: $DB_USER"
        echo ""
        if PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "SELECT 'Connected!' as status;" > /dev/null 2>&1; then
            echo "✅ Connection: OK"
            echo "📈 Tables:"
            PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "SELECT schemaname, tablename FROM pg_tables WHERE schemaname='public';" -t
        else
            echo "❌ Connection: FAILED"
        fi
        ;;
    *)
        echo "🐘 Conduit Database Utility"
        echo ""
        echo "Usage: $0 [command]"
        echo ""
        echo "Commands:"
        echo "  connect   - Connect to database interactively"
        echo "  tables    - List all tables"
        echo "  users     - Show users table info and data"
        echo "  status    - Show database connection status"
        echo "  reset     - Drop and recreate database (⚠️  DANGEROUS)"
        echo ""
        echo "Examples:"
        echo "  $0 connect"
        echo "  $0 users"
        echo "  $0 status"
        ;;
esac
