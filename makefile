setup:
	make proto
	cd src && go mod tidy

proto:
	protoc --go_out=./src --go-grpc_out=./src shared-protos/**/*.proto