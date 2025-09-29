# ![RealWorld Example App](logo.png)

> ### Go codebase containing real world examples (CRUD, auth, advanced patterns, etc) that adheres to the [RealWorld](https://github.com/gothinkster/realworld) spec and API.

### [RealWorld](https://github.com/gothinkster/realworld)

This codebase was created to demonstrate a fully fledged fullstack application built with **Go** including CRUD operations, authentication, routing, pagination, and more.

We've gone to great lengths to adhere to the **Go** community styleguides & best practices.

For more information on how this works with other frontends/backends, head over to the [RealWorld](https://github.com/gothinkster/realworld) repo.

## 🚀 **Features Implemented**

- ✅ **Authentication & Authorization** (JWT-based)
- ✅ **User Registration & Login**
- ✅ **Database Integration** (PostgreSQL with GORM)
- ✅ **RESTful API** following OpenAPI 3.0 specification
- ✅ **Clean Architecture** with domain-driven design
- ✅ **Environment Configuration** with .env support
- ✅ **Database Migrations** (Auto-migration with GORM)
- ✅ **Development Tools** (Database setup scripts)
- 🚧 **Articles CRUD** (Scaffold ready)
- 🚧 **Comments System** (Scaffold ready)
- 🚧 **User Profiles** (Scaffold ready)
- 🚧 **Follow/Unfollow Users** (Scaffold ready)
- 🚧 **Favorite Articles** (Scaffold ready)

## 🏗️ **Architecture & Technology Stack**

### **Technology Stack**
- **Go 1.21+** - Core language
- **Chi Router** - HTTP routing and middleware
- **GORM** - ORM for database operations
- **PostgreSQL** - Primary database
- **JWT** - Authentication tokens
- **OpenAPI 3.0** - API specification and code generation

### **Development Tools**
- **Air** - Hot-reload for development
- **Delve** - Go debugger with VS Code integration
- **oapi-codegen** - Code generation from OpenAPI specs
- **Custom Scripts** - Database management utilities

### **Project Structure**
```
├── cmd/                    # Application entrypoints
│   └── main.go            # Main application
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   ├── database/         # Database connection & migrations
│   ├── domain/           # Business logic & entities
│   │   └── user/        # User domain
│   ├── generated/        # OpenAPI generated code
│   ├── handlers/         # HTTP handlers (controllers)
│   ├── repository/       # Data access layer
│   └── services/         # Business services
├── pkg/                   # Public packages
│   ├── jwt/              # JWT utilities
│   └── logger/           # Logging utilities
├── scripts/              # Development scripts
│   ├── setup-db.sh      # Database setup
│   └── db.sh            # Database utilities
├── api/                  # API specifications
│   └── openapi.yml      # OpenAPI 3.0 spec
├── .vscode/               # VS Code configuration
│   └── launch.json       # Debug configurations
├── tmp/                  # Air temporary files
├── .air.toml            # Air configuration
├── .env                 # Environment variables (not in git)
├── example.env          # Environment template
├── Makefile             # Build and development commands
└── README.md            # This file
```

### **Design Patterns**
- **Clean Architecture** - Separation of concerns
- **Domain-Driven Design** - Business logic organization
- **Repository Pattern** - Data access abstraction
- **Dependency Injection** - Loose coupling
- **Composition over Inheritance** - Go idiomatic patterns

## 🚀 **Getting Started**

### **Prerequisites**
- Go 1.21 or higher
- PostgreSQL 12+
- Git
- Air (for hot-reload development) - `go install github.com/air-verse/air@latest`
- Delve (for debugging) - `go install github.com/go-delve/delve/cmd/dlv@latest`

### **Installation**

1. **Clone the repository**
   ```bash
   git clone https://github.com/joel-kores/Conduit_Real-World.git
   cd Conduit_Real-World
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Install development tools**
   ```bash
   # Install Air for hot-reload
   go install github.com/air-verse/air@latest

   # Install Delve for debugging
   go install github.com/go-delve/delve/cmd/dlv@latest
   ```

4. **Set up environment variables**
   ```bash
   cp example.env .env
   # Edit .env with your configuration
   ```

5. **Set up PostgreSQL database**
   ```bash
   # Automated setup (recommended)
   ./scripts/setup-db.sh

   # OR manual setup
   sudo -u postgres psql -c "CREATE USER conduituser WITH PASSWORD 'conduitpass';"
   sudo -u postgres psql -c "CREATE DATABASE conduitdb OWNER conduituser;"
   ```

6. **Start the application**
   ```bash
   # Development with hot-reload (recommended)
   make dev

   # OR production mode
   go run ./cmd/main.go
   ```

The API server will start on `http://localhost:8080`

### **Environment Configuration**

Copy `example.env` to `.env` and configure:

```bash
# Server Configuration
CONDUIT_SERVER_PORT=:8080
CONDUIT_LOG_LEVEL=debug

# JWT Configuration (⚠️ Change in production!)
CONDUIT_JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
CONDUIT_ACCESS_TOKEN_DURATION=15
CONDUIT_REFRESH_TOKEN_DURATION=1440

# Database Configuration
CONDUIT_DB_HOST=localhost
CONDUIT_DB_PORT=5432
CONDUIT_DB_USER=conduituser
CONDUIT_DB_PASSWORD=conduitpass
CONDUIT_DB_NAME=conduitdb
CONDUIT_DB_SSLMODE=disable
```

## 🛠️ **Development Tools**

### **Development Commands**
```bash
# Production build and run
make build              # Build binary
make run               # Build and run

# Development with hot-reload
make dev               # Start with auto-reload using Air

# Code generation
make gen               # Generate API code from OpenAPI spec

# Testing
make test              # Run all tests
```

### **Hot-Reload Development**
The project includes **Air** for automatic rebuilding and restarting during development:

```bash
# Start development server with hot-reload
make dev

# The server will automatically:
# - Watch for .go file changes
# - Rebuild the application
# - Restart the server
# - Show build errors in real-time
```

### **Debugging**
Full debugging support with VS Code and Delve:

```bash
# Debug with hot-reload
make debug-dev         # Start with debugger + hot-reload

# Debug production build
make debug             # Start with debugger only
```

**VS Code Debugging:**
1. Set breakpoints in your code (click left of line numbers)
2. Press `F5` or use Run > Start Debugging
3. Choose from available configurations:
   - **Launch Package** - Debug from VS Code directly
   - **Connect to Server** - Attach to running debug session
   - **Attach to Process** - Attach to any Go process

**Debug Controls:**
- `F5` - Continue/Start debugging
- `F10` - Step over
- `F11` - Step into
- `Shift+F11` - Step out
- `Shift+F5` - Stop debugging

### **Database Management**
```bash
# Check database status
./scripts/db.sh status

# List all tables
./scripts/db.sh tables

# Show users table and data
./scripts/db.sh users

# Connect to database interactively
./scripts/db.sh connect

# Reset database (⚠️ Warning: destructive)
./scripts/db.sh reset
```

### **API Testing**

**Register a new user:**
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "username": "testuser",
      "email": "test@test.com",
      "password": "password123"
    }
  }'
```

**Login:**
```bash
curl -X POST http://localhost:8080/api/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "email": "test@test.com",
      "password": "password123"
    }
  }'
```

## 🏃‍♂️ **How it works**

### **Request Flow**
1. **HTTP Request** → Chi Router
2. **Authentication Middleware** → JWT validation (for protected routes)
3. **Handler Layer** → Request validation & routing
4. **Service Layer** → Business logic execution
5. **Repository Layer** → Data persistence
6. **Database** → PostgreSQL with GORM

### **Authentication System**
- **JWT-based authentication** with configurable expiration
- **Password hashing** using bcrypt
- **Middleware protection** for authenticated routes
- **Token-based API access** following RealWorld spec

### **Database Design**
- **Auto-migration** on application startup
- **GORM ORM** for type-safe database operations
- **Connection pooling** for optimal performance
- **PostgreSQL** for production-grade reliability

## 📋 **API Endpoints**

### **Authentication**
- `POST /api/users` - Register user
- `POST /api/users/login` - Login user
- `GET /api/user` - Get current user (🚧)
- `PUT /api/user` - Update current user (🚧)

### **Articles** (🚧 Scaffold ready)
- `GET /api/articles` - List articles
- `POST /api/articles` - Create article
- `GET /api/articles/{slug}` - Get article
- `PUT /api/articles/{slug}` - Update article
- `DELETE /api/articles/{slug}` - Delete article

### **Profiles** (🚧 Scaffold ready)
- `GET /api/profiles/{username}` - Get profile
- `POST /api/profiles/{username}/follow` - Follow user
- `DELETE /api/profiles/{username}/follow` - Unfollow user

### **Comments** (🚧 Scaffold ready)
- `GET /api/articles/{slug}/comments` - Get comments
- `POST /api/articles/{slug}/comments` - Add comment
- `DELETE /api/articles/{slug}/comments/{id}` - Delete comment

## 👨‍💻 **Development Workflow**

### **Daily Development**
1. **Start development server:**
   ```bash
   make dev
   ```

2. **Make changes to your code** - The server automatically reloads

3. **Debug when needed:**
   ```bash
   # Start with debugger
   make debug-dev

   # In VS Code: F5 > "Connect to Server"
   ```

4. **Test your changes:**
   ```bash
   make test
   ```

5. **Check database:**
   ```bash
   ./scripts/db.sh status
   ./scripts/db.sh users
   ```

### **Adding New Features**
1. **Update OpenAPI spec** (`api/openapi.yml`)
2. **Regenerate code:**
   ```bash
   make gen
   ```
3. **Implement handlers, services, and repositories**
4. **Test your implementation**
5. **Update documentation**

### **Common Development Tasks**
```bash
# View logs in real-time
tail -f tmp/build-errors.log

# Reset database for testing
./scripts/db.sh reset

# Generate new API code after spec changes
make gen

# Run specific tests
go test ./internal/services/...
```

## 🧪 **Testing**

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Test database connection
./scripts/db.sh status
```

## 🚀 **Deployment**

### **Production Checklist**
- [ ] Update `CONDUIT_JWT_SECRET` with a secure random key
- [ ] Set `CONDUIT_LOG_LEVEL=info` or `warn`
- [ ] Configure production database credentials
- [ ] Set `CONDUIT_DB_SSLMODE=require` for secure connections
- [ ] Configure proper CORS settings
- [ ] Set up reverse proxy (nginx/traefik)
- [ ] Configure TLS/SSL certificates

### **Docker Support** (Coming Soon)
```bash
# Build and run with Docker
docker-compose up -d
```

## 🤝 **Contributing**

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📚 **Learning Resources**

- [RealWorld Specification](https://github.com/gothinkster/realworld/tree/main/api)
- [Go Documentation](https://golang.org/doc/)
- [Chi Router](https://github.com/go-chi/chi)
- [GORM Documentation](https://gorm.io/docs/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)

## 📄 **License**

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙋‍♂️ **Support**

- **Issues**: [GitHub Issues](https://github.com/joel-kores/Conduit_Real-World/issues)
- **Discussions**: [GitHub Discussions](https://github.com/joel-kores/Conduit_Real-World/discussions)
- **RealWorld Community**: [RealWorld Spec](https://github.com/gothinkster/realworld)

---

**Built with ❤️ using Go and following RealWorld specifications**

