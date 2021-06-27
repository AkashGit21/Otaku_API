
clean:
	rm -r pb/*

client: 
	go run cmd/client/main.go -address 127.0.0.1:8080

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root otaku_db

dropdb:
	docker exec -it postgres13 dropdb otaku_db

gen: 
	protoc --proto_path=proto proto/*.proto  --go_out=pb --go-grpc_out=pb --grpc-gateway_out=pb --openapiv2_out=:swagger

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/otaku_db?sslmode=disable" -verbose down

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/otaku_db?sslmode=disable" -verbose up

postgresconnect:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:13-alpine

server: 
	go run cmd/server/main.go	-port 50051

rest:
	go run cmd/server/main.go -port 8081 -type rest -endpoint 127.0.0.1:8080

sqlc:
	sqlc generate

.PHONY: gen clean server rest client postgresconnect createdb dropdb migrateup migratedown sqlc