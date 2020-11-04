# Usage
# make proto

proto:
	@protoc -I. \
	  -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	  -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
	  --proto_path=.:${GOPATH}/src \
	  --go_out=plugins=grpc,paths=source_relative:. api/proto/*.proto; \

	@echo ✓ protobuf compiled; \
