# Inventory Reservation Lab

A Go learning project for an inventory reservation API backed by PostgreSQL.

Currently, the server exposes a JSON health endpoint. Product input validation and a product-table migration are available. The API is not connected to PostgreSQL yet.

## Run

```powershell
go run ./cmd/api
```

The server listens on `http://127.0.0.1:8080`. Stop it with Ctrl+C.

## Check the server

```powershell
curl.exe -i http://127.0.0.1:8080/health
```

Expected response: `200 OK` with `Content-Type: application/json` and a body of `{"status":"ok"}`.

## Test

```powershell
go test ./...
```

## Local database setup

Install PostgreSQL 18 with Command Line Tools.

Connect as the administrator:

```powershell
& "C:\Program Files\PostgreSQL\18\bin\psql.exe" -h 127.0.0.1 -p 5432 -U postgres -d postgres -W
```

Create the project login and database:

```sql
CREATE ROLE inventory_app LOGIN;
```

Set the login password interactively:

```text
\password inventory_app
```

Then create the database:

```sql
CREATE DATABASE inventory_reservation_lab OWNER inventory_app;
```

Exit psql with `\q`.

From the repository root, apply the migration once to the new database:

```powershell
& "C:\Program Files\PostgreSQL\18\bin\psql.exe" -h 127.0.0.1 -p 5432 -U inventory_app -d inventory_reservation_lab -W -v ON_ERROR_STOP=1 --single-transaction -f migrations/001_create_products.sql
```

The SQL files are currently applied manually; there is no migration tracking yet.
