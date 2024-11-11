# go-grpc-unary-operation
go-grpc-unary-operation

This is simple application showing imlementation of Unary GRPC implementation


//Make sure below requirements are installed

1. brew install protobuf-c (to verify use: protoc --version)

2. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    protoc-gen-go --version
    
3. to generate grpc models use commands: protoc --go-grpc_out=.  --go_out=.  *.proto

// protoc : will direct tot proto compiler
// --go-grpc_out=. : with will tells grpc to generate files in the same folder
// --go_out=. : generated moesl should be in the same paths
// *.proto : henerate for all the protot paths


Steps of Simualtion
1. run server: go run server.go (inside server folder)
2. run client: go run client.go (inside client folder)