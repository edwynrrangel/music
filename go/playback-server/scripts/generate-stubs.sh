#!/bin/bash

for proto in ./proto/*.proto; do
  DIR_NAME=$(basename "$proto" .proto)
  if [ ! -d "./internal/adapter/grpc/$DIR_NAME" ]; then
    mkdir -p "./internal/adapter/grpc/$DIR_NAME"
  fi
  protoc -I=proto --go_out=./internal/adapter/grpc/ \
    --go-grpc_out=./internal/adapter/grpc/ \
    "$proto"
done
