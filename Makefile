start:
	@go run main.go

start_client:
	@go run client/main.go
	
gen:
	@protoc --proto_path=proto proto/crypto/crypto.proto --go_out=plugins=grpc:proto

seed:
	@go run database/seeder/seeder.go