#!/bin/bash

protoc \
  --go_out=./coffee_Shop_proto --go_opt=paths=source_relative \
  --go-grpc_out=./coffee_Shop_proto --go-grpc_opt=paths=source_relative \
  coffee_shop.proto
