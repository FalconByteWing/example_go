# Go Project Structure Example

This repository demonstrates a practical example of structuring a Go microservice project, using a payout service as an example case.

## Project Overview

This example showcases:
- project structure for Go microservices
- Database management with migrations
- Type-safe SQL with SQLC
- API documentation with Swagger
- Makefile for common development tasks
- Docker setup for development

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose
- Make (optional, but recommended)

## Getting Started

1. Clone the repository:

2. Set up environment variables:

3. Initialize the project:
```bash
# Initialize database and generate required code
make init-db        # Creates and initializes the database
go mod tidy        # Download and tidy Go dependencies
make sqlc          # Generate type-safe SQL code
make swagger-tsp-service-payout  # Generate Swagger documentation
make migrate-up    # Run database migrations
```

4. Start the server:
```bash
make run-tsp-service-payout
```

5. Access Swagger Documentation:
```bash
make swagger-serve-tsp-service-payout  # Serves Swagger UI locally
```

## Development Commands

- `make build`: Build the application
- `make test`: Run tests
- `make lint`: Run linters
- `make init-db`: Initialize the database with required schemas
- `make sqlc`: Generate type-safe SQL code using SQLC
- `make swagger-tsp-service-payout`: Generate Swagger documentation
- `make swagger-serve-tsp-service-payout`: Serve Swagger UI locally
- `make generate`: Generate code (SQLC, Swagger)
- `make down`: Stop Docker containers
- `make migrate-up`: Apply database migrations
- `make migrate-down`: Rollback migrations

## API Documentation

Once the server is running, access the Swagger documentation at:
```
http://localhost:8080/swagger/index.html
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.




