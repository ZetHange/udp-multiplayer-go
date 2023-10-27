run server:
	go run ./cmd/server/main.go
run client:
	go run ./cmd/client/client.go
build:
	go build -o server -ldflags "-s -w" ./cmd/server/main.go
grpc-build:	
	protoc --go_out=. --go-grpc_out=. proto/multiplayer.proto
