## Setup

Run Initiate
```console
go mod tidy
```

Run Migrations

1. go run cmd/bun/main.go db init
    Create file Migration
    go run cmd/bun/main.go db create_go create_orders
2. update migration file with models
3. go run cmd/bun/main.go db migrate
4. go run cmd/bun/main.go db rollback