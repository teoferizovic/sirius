# Golang microservice - grpc example
## command to execute protc files
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user.proto
## Start Server
go run server/*
## Start Client
cd client && go run main.go