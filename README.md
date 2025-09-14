# Demo Gin API

A standard Go backend service using Gin, PostgreSQL, sqlc, and OpenAPI.

## Project Structure

```
demo-gin/
├── api/                    # OpenAPI specifications
│   └── openapi.yaml       # OpenAPI 3.0 specification
├── cmd/                   # Application entry points
│   └── server/
│       └── main.go        # Main application
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   ├── db/               # Database related code
│   │   ├── queries/      # SQL queries for sqlc
│   │   └── sqlc/         # Generated code by sqlc
│   ├── handlers/         # HTTP request handlers
│   ├── middleware/       # HTTP middleware
│   ├── models/           # Data models
│   ├── services/         # Business logic
│   └── utils/            # Utility functions
├── migrations/            # Database migrations
├── pkg/                   # Public packages
│   ├── database/         # Database connection
│   └── logger/           # Logging utilities
├── docs/                  # Generated Swagger docs
├── .env.example          # Environment variables example
├── .gitignore
├── go.mod
├── go.sum
├── Makefile              # Build and deployment scripts
├── sqlc.yaml             # sqlc configuration
└── README.md

```

## Prerequisites

- Go 1.22+
- PostgreSQL 14+
- [sqlc](https://sqlc.dev/)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [swag](https://github.com/swaggo/swag) (for Swagger docs)

## Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd demo-gin
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Setup environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

4. **Run database migrations**
   ```bash
   make migrate-up
   ```

5. **Generate sqlc code**
   ```bash
   make sqlc
   ```

6. **Generate Swagger documentation**
   ```bash
   make swagger
   ```

## Development

### Run the server
```bash
make run
```

### Build the application
```bash
make build
```

### Run tests
```bash
make test
```

## API Documentation

- **OpenAPI Specification**: `api/openapi.yaml`
- **Swagger UI**: After running the server, visit `http://localhost:8080/swagger/index.html`

## Available Endpoints

### Authentication
- `POST /api/v1/auth/register` - Register a new user
- `POST /api/v1/auth/login` - Login user

### Users (Protected)
- `GET /api/v1/users` - List users
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Posts
- `GET /api/v1/posts` - List posts (public)
- `GET /api/v1/posts/:id` - Get post by ID (public)
- `POST /api/v1/posts` - Create post (protected)
- `PUT /api/v1/posts/:id` - Update post (protected)
- `DELETE /api/v1/posts/:id` - Delete post (protected)

### Health
- `GET /api/v1/health` - Health check

## Database Schema

The application includes two main tables:
- **users**: User accounts with authentication
- **posts**: Content posts linked to users

See `migrations/000001_init_schema.up.sql` for the complete schema.

## Makefile Commands

```bash
make help        # Show available commands
make run         # Run the application
make build       # Build the application
make test        # Run tests
make migrate-up  # Run database migrations up
make migrate-down # Run database migrations down
make sqlc        # Generate code from SQL
make swagger     # Generate swagger documentation
make deps        # Download dependencies
make clean       # Clean build artifacts
```

## TODO

- [ ] Implement JWT authentication
- [ ] Add password hashing (bcrypt)
- [ ] Complete sqlc integration in handlers
- [ ] Add input validation
- [ ] Implement proper error handling
- [ ] Add logging with structured logs
- [ ] Add unit and integration tests
- [ ] Add rate limiting
- [ ] Add request/response logging middleware
- [ ] Add graceful shutdown
- [ ] Add Docker support
- [ ] Add CI/CD pipeline