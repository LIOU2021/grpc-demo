echo "Try to generate pb.go"
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative users.proto
echo "Finished to generate pb.go"
