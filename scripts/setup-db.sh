#!/bin/bash

# Conduit RealWorld Database Setup Script
# This script sets up PostgreSQL database for the Conduit project

set -e

echo "🐘 Setting up PostgreSQL database for Conduit RealWorld..."

# Load environment variables
source .env 2>/dev/null || echo "Warning: .env file not found, using defaults"

# Extract database settings from environment variables or use defaults
DB_USER=${CONDUIT_DB_USER:-conduituser}
DB_PASSWORD=${CONDUIT_DB_PASSWORD:-conduitpass}
DB_NAME=${CONDUIT_DB_NAME:-conduitdb}
DB_HOST=${CONDUIT_DB_HOST:-localhost}
DB_PORT=${CONDUIT_DB_PORT:-5432}

echo "📊 Database Configuration:"
echo "  Host: $DB_HOST:$DB_PORT"
echo "  Database: $DB_NAME"
echo "  User: $DB_USER"
echo ""

# Check if PostgreSQL is running
if ! sudo systemctl is-active --quiet postgresql; then
    echo "❌ PostgreSQL is not running. Starting it..."
    sudo systemctl start postgresql
    echo "✅ PostgreSQL started"
fi

# Create user if it doesn't exist
echo "👤 Creating database user '$DB_USER'..."
sudo -u postgres psql -c "CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';" 2>/dev/null || echo "User already exists"

# Create database if it doesn't exist
echo "🗄️  Creating database '$DB_NAME'..."
sudo -u postgres psql -c "CREATE DATABASE $DB_NAME OWNER $DB_USER;" 2>/dev/null || echo "Database already exists"

# Grant privileges
echo "🔐 Granting privileges..."
sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;"
sudo -u postgres psql -c "ALTER USER $DB_USER CREATEDB;"

# Test connection
echo "🔍 Testing database connection..."
if PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "SELECT 'Connection successful!' as status;" > /dev/null 2>&1; then
    echo "✅ Database connection successful!"
else
    echo "❌ Database connection failed!"
    exit 1
fi

echo ""
echo "🎉 Database setup complete!"
echo ""
echo "📝 Next steps:"
echo "  1. Run your Go application: go run ./cmd/main.go"
echo "  2. The app will auto-migrate tables on startup"
echo "  3. Test endpoints with curl or Postman"
echo ""
echo "🔧 Useful commands:"
echo "  Connect to DB: PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME"
echo "  Check tables: \\dt"
echo "  Drop DB: sudo -u postgres psql -c \"DROP DATABASE $DB_NAME;\""
echo ""
