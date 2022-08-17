## Setup

GO version 1.17



Run first
```console
go mod tidy
```

#Create Database on your local, use MySQL 

#Set .env value MYSQL_URL 

#Run Migrations
```
1. go run cmd/bun/main.go db init
2. go run cmd/bun/main.go db migrate
```

#Start Service
```> make run-grpc
go run grpc-app/main.go
ts=2022-08-17T14:02:21.922+0700 level=info msg="GRPC service started on port :50051"
```
