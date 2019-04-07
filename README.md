# Pluralsight tutorial:  Enhancing Application Communication with gRPC

## Generating Source Code
```
protoc -I ./pb ./pb/messages.proto --go_out=plugins=grpc:./src
```