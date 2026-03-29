# GO TODO API

A production-ready RESTful API built with Go for managing todos with user authentication. Features JWT-based authentication, PostgreSQL database, and full Docker support.

## Features

- **User Authentication**
  - User registration with secure password hashing (bcrypt)
  - Login with JWT token generation
  - HttpOnly cookie-based token storage for enhanced security

- **Todo Management**
  - Create todos linked to authenticated users
  - User-scoped data access

- **Security**
  - JWT-based authentication middleware
  - Password hashing with bcrypt
  - Secure cookie handling
  - User-scoped database operations

- **Infrastructure**
  - Dockerized application with multi-stage builds
  - PostgreSQL database with connection pooling
  - Database migrations for schema versioning
  - Live reload for development (Air)
  - Health checks for monitoring

## Tech Stack

- **Language**: Go 1.25+
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin) - HTTP router and middleware
- **Database**: PostgreSQL 16 with [pgx/v5](https://github.com/jackc/pgx) driver
- **Authentication**: [golang-jwt/jwt](https://github.com/golang-jwt/jwt) for JWT tokens
- **Password Hashing**: bcrypt
- **Configuration**: [godotenv](https://github.com/joho/godotenv)
- **Development**: [Air](https://github.com/cosmtrek/air) for live reload
- **Containerization**: Docker & Docker Compose

## Prerequisites

- **Go** 1.25 or higher
- **PostgreSQL** (or use Docker)
- **Docker & Docker Compose** (optional, but recommended)

## Installation & Setup

### 1. Clone the Repository

```bash
git clone <repository-url>
cd GO_TODO_API
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Configure Environment Variables

Create a `.env` file in the project root:

```env
DATABASE_URI=postgres://user:password@localhost:5432/tododb?sslmode=disable
PORT=9090
JWT_SECRET=your-secret-key-here-change-in-production
```

## Running the Application

### Option 1: Using Docker (Recommended)

The easiest way to run the application with all dependencies:

```bash
# Start all services (API + PostgreSQL)
docker-compose up --build

# Run in detached mode
docker-compose up -d

# View logs
docker-compose logs -f api

# Stop services
docker-compose down
```

The API will be available at `http://localhost:9090`

### Option 2: Local Development

#### Step 1: Start PostgreSQL

Make sure PostgreSQL is running locally, or start it via Docker:

```bash
docker run -d \
  --name postgres-todo \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=tododb \
  -p 5432:5432 \
  postgres:16-alpine
```

#### Step 2: Run Database Migrations

```bash
chmod +x scripts/migrate.sh
./scripts/migrate.sh up
```

#### Step 3: Start the Server

**With Air (live reload):**
```bash
air
```

**Without Air:**
```bash
go run cmd/api/main.go
```

**Build and run:**
```bash
go build -o bin/api cmd/api/main.go
./bin/api
```

## Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `DATABASE_URI` | PostgreSQL connection string | `postgres://user:password@localhost:5432/tododb?sslmode=disable` |
| `PORT` | Server port | `9090` |
| `JWT_SECRET` | Secret key for JWT signing | `your-super-secret-key` |

## Database Migrations

The project uses a custom migration script for database schema management.

### Migration Commands

```bash
# Apply all pending migrations
./scripts/migrate.sh up

# Rollback the last migration
./scripts/migrate.sh down

# Create a new migration
./scripts/migrate.sh create <migration_name>

# Reset database (drop all and recreate)
./scripts/migrate.sh reset

# Show help
./scripts/migrate.sh
```

### Migration Files

Migration files are located in the `migrations/` directory:
- `001_create_todos_table.sql`
- `002_create_users_table.sql`
- `003_add_user_id_to_todos.sql`
- `004_add_name_to_users.sql`

## API Endpoints

### Public Endpoints

#### Health Check
```http
GET /
```

**Response:**
```json
{
  "message": "SUCCESS!!"
}
```

#### Register User
```http
POST /api/v1/register
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword123"
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2026-03-29T10:30:00Z"
}
```

**Example with curl:**
```bash
curl -X POST http://localhost:9090/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword123"
  }'
```

#### Login
```http
POST /api/v1/login
Content-Type: application/json
```

**Request Body:**
```json
{
  "email": "john@example.com",
  "password": "securepassword123"
}
```

**Response (200 OK):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

Sets an HttpOnly cookie named `token` that expires in 24 hours.

**Example with curl:**
```bash
curl -X POST http://localhost:9090/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "securepassword123"
  }' \
  -c cookies.txt  # Save cookies to file
```

### Protected Endpoints

These endpoints require authentication. Include the JWT token in the `token` cookie.

#### Create Todo
```http
POST /api/v1/todos
Content-Type: application/json
Cookie: token=<jwt-token>
```

**Request Body:**
```json
{
  "title": "Complete project documentation"
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "user_id": 1,
  "title": "Complete project documentation",
  "completed": false,
  "created_at": "2026-03-29T10:35:00Z",
  "updated_at": "2026-03-29T10:35:00Z"
}
```

**Example with curl:**
```bash
curl -X POST http://localhost:9090/api/v1/todos \
  -H "Content-Type: application/json" \
  -b cookies.txt \  # Use cookies from login
  -d '{
    "title": "Complete project documentation"
  }'
```

## Project Structure

```
GO_TODO_API/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── auth/
│   │   ├── jwt.go                  # JWT token generation & validation
│   │   └── passcode.go             # Password hashing & comparison
│   ├── config/
│   │   └── config.go               # Environment configuration loader
│   ├── database/
│   │   └── postgres.go             # PostgreSQL connection setup
│   ├── handlers/
│   │   ├── auth.handler.go         # Register & login handlers
│   │   └── todo.handler.go         # Todo CRUD handlers
│   ├── middleware/
│   │   └── auth.middleware.go      # JWT authentication middleware
│   ├── models/
│   │   ├── user.go                 # User database model
│   │   ├── todos.go                # Todo database model
│   │   ├── request.go              # API request DTOs
│   │   └── response.go             # API response DTOs
│   └── repository/
│       ├── user.repository.go      # User database operations
│       └── todos.repository.go     # Todo database operations
├── migrations/                      # SQL migration files
├── scripts/
│   └── migrate.sh                  # Migration management script
├── .air.toml                       # Air live reload configuration
├── .env                            # Environment variables
├── docker-compose.yml              # Multi-container orchestration
├── Dockerfile                      # Application container definition
├── go.mod                          # Go module dependencies
└── go.sum                          # Dependency checksums
```

## Development

### Live Reload with Air

The project is configured with Air for automatic reloading during development:

```bash
# Make sure Air is installed
go install github.com/cosmtrek/air@latest

# Run with live reload
air
```

Air will watch for changes in `.go` files and automatically rebuild and restart the server.

### Building for Production

```bash
# Build binary
go build -o bin/api cmd/api/main.go

# Run binary
./bin/api

# Build with optimizations
go build -ldflags="-s -w" -o bin/api cmd/api/main.go
```

### Using Docker in Development

```bash
# Rebuild after code changes
docker-compose up --build

# View logs
docker-compose logs -f api

# Access PostgreSQL
docker-compose exec db psql -U user -d tododb

# Run migrations in container
docker-compose exec api ./scripts/migrate.sh up
```

## Docker Configuration

### Services

- **api**: Go application (port 9090)
- **db**: PostgreSQL 16 (port 5432)

### Features

- Multi-stage Docker build for optimized image size (~34 MB)
- Health checks for both services
- Automatic restarts
- Persistent volume for database data
- Non-root user execution for security

### Exposed Ports

- **9090**: API server
- **5432**: PostgreSQL database

## Database Schema

### Users Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Todos Table
```sql
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Architecture

This project follows a clean architecture pattern with clear separation of concerns:

- **cmd/**: Application entry points
- **internal/**: Private application code
  - **handlers**: HTTP request handlers (presentation layer)
  - **repository**: Data access layer
  - **models**: Domain models and DTOs
  - **auth**: Authentication logic
  - **middleware**: HTTP middleware components
  - **config**: Configuration management
  - **database**: Database connection setup

## Future Enhancements

- [ ] Complete CRUD operations for todos (GET, UPDATE, DELETE)
- [ ] List todos with pagination and filtering
- [ ] Mark todos as completed/uncompleted
- [ ] Automated test suite (unit and integration tests)
- [ ] API documentation with Swagger/OpenAPI
- [ ] Rate limiting and request throttling
- [ ] Refresh token mechanism
- [ ] User profile management
- [ ] Todo categories/tags
- [ ] Due dates and reminders
- [ ] Logging middleware

## Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the [MIT License](LICENSE).

---

**Note**: This project is currently in development. The todo management feature currently supports creation only. Additional CRUD operations and features are planned for future releases.
