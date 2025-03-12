# Go Fiber Postgres Project

This project is a simple REST API built with Go, Fiber, and PostgreSQL. It includes basic CRUD operations for managing users.

## Prerequisites

- Go 1.18+
- Docker (for running PostgreSQL)
- Make sure you have the necessary environment variables set up in a `.env` file.

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/supacheep-first/go-fiber-postgres.git
cd go-fiber-postgres
```

### Setup Environment Variables

Create a `.env` file in the root directory with the following content:

```properties
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgresuser
DB_PASSWORD=postgrespassword
DB_NAME=fiberdb
SSL_MODE=disable
```

### Run PostgreSQL with Docker

```sh
docker-compose up -d
```

### Install Dependencies

```sh
go mod tidy
```

### Run the Application

```sh
go run main.go
```

The application will be running on `http://localhost:3000`.

### Running Tests

To run the tests, use the following command:

```sh
go test ./tests
```

## Project Structure

- `main.go`: Entry point of the application.
- `database/`: Contains database connection and migration logic.
- `models/`: Contains the data models.
- `routes/`: Contains the route definitions.
- `handlers/`: Contains the request handlers.
- `validators/`: Contains the validation logic.
- `tests/`: Contains the test cases.

## Adding Features

To add new features, follow these steps:

1. **Define the Model**: Add a new model in the `models` package.
2. **Create Migrations**: Update the `database` package to include migrations for the new model.
3. **Add Routes**: Define new routes in the `routes` package.
4. **Implement Handlers**: Implement the request handlers in the `handlers` package.
5. **Write Tests**: Add test cases in the `tests` package to ensure the new feature works as expected.

## Contributing

Feel free to open issues or submit pull requests for any improvements or new features.
