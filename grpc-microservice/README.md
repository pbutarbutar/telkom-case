## Setup

GO version 1.17



Run first
```console
go mod tidy
```

#Create Database on your local, use MySQL 

#Set .env value MYSQL_URL 

#Run Migrations

1. go run cmd/bun/main.go db init
2. go run cmd/bun/main.go db migrate
