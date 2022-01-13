generate:
	@protoc --proto_path=proto proto/crypto/crypto.proto --go_out=plugins=grpc:proto

run:
	@echo "---- Running Server ----"
	@go run main.go

run_client:
	@echo "---- Running Client ----"
	@go run client/main.go