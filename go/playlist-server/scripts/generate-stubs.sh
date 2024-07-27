#!/bin/bash

for proto in ./proto/*.proto; do
  DIR_NAME=$(basename "$proto" .proto)
  if [ ! -d "./internal/port/grpc/$DIR_NAME" ]; then
    mkdir -p "./internal/port/grpc/$DIR_NAME"
  fi
  protoc --go_out=./internal/port/grpc/ \
    --go-grpc_out=./internal/port/grpc/ \
    "$proto"
done
