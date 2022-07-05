setup:
	make proto
	cd src && go mod tidy

proto:
	protoc --go_out=./src --go-grpc_out=./src shared-protos/**/*.proto

dev:
	go install github.com/cespare/reflex@latest
	reflex -r "\.go$\" -s -- sh -c "cd src && go run ."
