#!/bin/bash

# target project
target=user

# rpc generate
goctl rpc protoc ./${target}/*.proto --go_out=./${target}/pb --go-grpc_out=./${target}/pb  --zrpc_out=./${target} --style=goZero --home ../deploy/goctl/1.5.0