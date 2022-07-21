#!/usr/bin/env sh

# Install proto3 from brew
# brew install protobuf
#
# update protoc Go bindings via
# go get -u google.golang.org/protobuf/cmd/protoc-gen-go
# go install google.golang.org/protobuf/cmd/protoc-gen-go
#
# go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
# See also
# https://github.com/grpc/grpc-go/tree/master/examples

protoc addsvc.proto --go_out=. --go-grpc_out=.