# Wordiest API

Run API:

```bash
go run ./cmd/server
```

Create Swagger docs:

```bash
swag init -g cmd/server/main.go
```

Build server:

```bash
go build -o bin/server ./cmd/server
```
