# kimbap
Hobby project for learning Go and gRPC

[![CircleCI](https://circleci.com/gh/stopa323/kimbap.svg?style=svg)](https://circleci.com/gh/stopa323/kimbap)

# Protobuf generation
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative bmc.proto
```
