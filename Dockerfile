
# TODO

# Run the app
FROM golang:1.16.5-alpine3.13

WORKDIR /app

COPY . .

RUN go run cmd/server/main.go -port 8081 -type rest -endpoint 127.0.0.1:8080

EXPOSE 8081