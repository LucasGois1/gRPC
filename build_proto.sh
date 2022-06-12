echo "Building proto files and gRPC connection entities"
protoc --proto_path=protobuf/ protobuf/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.
echo "Done!"
