# Go-Vite

This is a simple web server written in Go that uses SQLite as the database.
It serves a React + TypeScript + Vite frontend.
It is designed to be a minimal example of how to set up a web server with a frontend and a database.

## Features

- React + TypeScript + Vite (embedded in Go)
- SQLite database with schema migrations
- REST API endpoints
- Docker support
- Environment variable configuration (using `.env` file)

## Running the server

Build the frontend:

```shell
cd frontend
bun install
bun run build
```

Then build the Go server:

```shell
go build -o app .
```

Or if you want to run the server in development mode, you can use:

```shell
go run main.go
```

The server will start on `localhost:8080` by default.
You can change the port by setting the `PORT` environment variable.

## Docker

To start the server in a Docker container, you can use the provided `docker-compose.yml` file.
Make sure you have Docker and Docker Compose installed.
To build and run the Docker container, use the following command:

```shell
docker-compose up --build
```

This will build the Docker image and start the container.
The server will be accessible at `http://localhost:8080`.
To stop the container, use:

```shell
docker-compose down
```

## Environment Variables

The server uses environment variables to configure the database connection and other settings.
You can create a `.env` file in the root directory of the project to set these variables.
The following environment variables are supported:

- `PORT`: The port on which the server will listen. Default is `8080`.

## Database

Uses the [migrate](github.com/golang-migrate/migrate/v4) library to manage database migrations.

### Database Migrations

Creating new migrations:

```shell
migrate create -ext sql -dir data/migrations -seq create_users_table
```

Update the created migration file with the SQL statements.

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);
```

The migrations are applied automatically when the server starts.

To manually run the migrations:

```shell
# Run the migrations
migrate -path data/migrations -database "sqlite3://./data/go-vite.db" up

# Rollback the migrations
migrate -path data/migrations -database "sqlite3://./data/go-vite.db" down

# Rollback the last migration
migrate -path data/migrations -database "sqlite3://./data/go-vite.db" down 1
```
