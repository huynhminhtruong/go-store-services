# plugin protoc-gen-go helps protoc(Protocol Buffers Compiler) creates Go files from .proto
install-protobuf:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# plugin protoc-gen-go-grpc helps protoc(Protocol Buffers Compiler) creates gRPC services in Go from .proto
# includes interfaces and methods
install-grpc:
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

mkdir-service:
	@mkdir -p ./src/services/${SERVICE_NAME}

# need install protoc(Protocol Buffers Compiler) first
build-service:
	@protoc -I ./proto/${SERVICE_NAME} \
		--go_out ./src/services/${SERVICE_NAME} \
		--go_opt paths=source_relative \
		--go-grpc_out ./src/services/${SERVICE_NAME} \
		--go-grpc_opt paths=source_relative \
		./proto/${SERVICE_NAME}/*.proto