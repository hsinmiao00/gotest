# Protobuf test

build: protoc -I $GOPATH/src:. --go_out=. test.proto

https://gowalker.org/github.com/golang/protobuf/proto
