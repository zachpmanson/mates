# Backend (Go) — mates

This directory contains a minimal Go HTTP server scaffold.

Available endpoints:

- GET /health — returns JSON with status and uptime
- GET /api/reports — returns a sample JSON array of reports

Run locally:

```sh
# from repository root
cd backend
go run main.go
```

Build:

```sh
cd backend
go build -o mates-server
./mates-server
```

Run tests:

```sh
cd backend
go test ./...
```

Environment variables:

- PORT — optional, defaults to 8080
