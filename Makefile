protobuf: pb/service.proto
	protoc --go_out=. --go-grpc_out=. pb/service.proto
