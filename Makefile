run_tests:
	go test server/...

generate_proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative server/proto/health/health.proto