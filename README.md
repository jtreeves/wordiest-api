# Wordiest API

Run API:

```bash
go run ./cmd/api
```

Create Swagger docs:

```bash
swag init -g cmd/api/main.go
```

Build API:

```bash
go build -o bin/api ./cmd/api
```
