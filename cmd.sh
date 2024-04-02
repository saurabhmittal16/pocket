#!/bin/bash

# Define your methods
genPlaygroundRPC() {
    protoc \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    playground/hello.proto
}

genServiceRPC() {
    protoc \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    service/service.proto
}

genRPC() {
    protoc \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    $1
}

getPID() {
    lsof -i :$1
}

# Check the argument passed and call the corresponding method
if [ "$1" = "genPlaygroundRPC" ]; then
    genPlaygroundRPC
elif [ "$1" = "genServiceRPC" ]; then
    genServiceRPC
elif [ "$1" = "genRPC" ]; then
    genRPC $2
elif [ "$1" = "getPID" ]; then
    getPID $2
else
    echo "Usage: $0 [genPlaygroundRPC | genServiceRPC | genRPC <.proto>]"
    exit 1
fi