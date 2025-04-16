# nanites

## Running the server

```shell
go run main.go
```

The server will start on `localhost:8080` by default.
You can change the port by setting the `PORT` environment variable.

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
migrate -path data/migrations -database "sqlite3://./data/nanites.db" up

# Rollback the migrations
migrate -path data/migrations -database "sqlite3://./data/nanites.db" down

# Rollback the last migration
migrate -path data/migrations -database "sqlite3://./data/nanites.db" down 1
```
