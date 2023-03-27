#!/bin/bash
# Builds the protocol buffer files Locally 

protoc  --go_out=. \
        --go_opt=paths=source_relative \
        --go-grpc_out=. \
        --go-grpc_opt=paths=source_relative \
        ./*.proto