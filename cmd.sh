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

# Check the argument passed and call the corresponding method
if [ "$1" = "genPlaygroundRPC" ]; then
    genPlaygroundRPC
elif [ "$1" = "genServiceRPC" ]; then
    genServiceRPC
else
    echo "Usage: $0 [genPlaygroundRPC|genServiceRPC]"
    exit 1
fi