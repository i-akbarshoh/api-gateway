proto-gen:
	protoc --go_out=./internal/pkg --go_opt=paths=source_relative \
	--go-grpc_out=./internal/pkg --go-grpc_opt=paths=source_relative \
	proto/auth-service.proto

run:
	go run cmd/main.go