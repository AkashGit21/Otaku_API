
gen: 
	protoc --proto_path=proto proto/*.proto  --go_out=pb --go-grpc_out=pb --grpc-gateway_out=pb --openapiv2_out=:swagger

clean:
	rm -r pb/*

server: 
	go run cmd/server/main.go	