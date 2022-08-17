create:
	protoc --proto_path=protos protos/health.proto --go_out=grpc-app/protos
	protoc --proto_path=protos protos/health.proto --go-grpc_out=grpc-app/protos

	protoc --proto_path=protos protos/carts/carts.proto --go_out=grpc-app/protos/carts
	protoc --proto_path=protos protos/carts/carts.proto --go-grpc_out=grpc-app/protos/carts

mock:
	mockgen -package=mock_repo -source=shared/repository/carts/carts.go -destination=shared/repository/carts/mocks/carts_repo_mock.go
	mockgen -package=mock_usecase -source=shared/usecase/carts/carts.go -destination=shared/usecase/carts/mocks/carts_usecase_mock.go
		

clean:
	rm grpc-app/protos/*/*.go

test:
	go test ./... -coverprofile coverage.out
	rm -rf coverage.out

injector:
	wire ./grpc-app/service/injector/health
	wire ./grpc-app/service/injector/carts	

run-grpc:
	go run grpc-app/main.go

migrate:
	go run cmd/bun/main.go db migrate

.PHONY: create clean mock test injector run 