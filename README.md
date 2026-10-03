# Inventory Reservation Lab

A Go learning project for an inventory reservation API backed by PostgreSQL.

Currently, the server exposes a health endpoint. Inventory and database
functionality are not implemented yet.

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
