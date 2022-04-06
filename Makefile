run_tests:
	go test ./...

generate_proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative server/proto/health/health.proto

start_client:
	go run client/main.go

start_server:
	go run server/main.go