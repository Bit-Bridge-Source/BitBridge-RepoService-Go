module github.com/Bit-Bridge-Source/BitBridge-RepoService-Go

go 1.21.3

require (
	github.com/Bit-Bridge-Source/BitBridge-CommonService-Go v1.10.4
	go.mongodb.org/mongo-driver v1.12.1
)

require (
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.3.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/Bit-Bridge-Source/BitBridge-CommonService-Go => ../common-service
