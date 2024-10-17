protoc --go_out=../go-grpc/hoarddata --go-grpc_out=../go-grpc/hoarddata --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative hoard-data.proto
protoc --go_out=../go-grpc/hoardclient --go-grpc_out=../go-grpc/hoardclient --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative hoard-client.proto
protoc --go_out=../go-grpc/hoardauth --go-grpc_out=../go-grpc/hoardauth --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative hoard-auth.proto
protoc --go_out=../go-grpc/hoardcompute --go-grpc_out=../go-grpc/hoardcompute --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative hoard-compute.proto
protoc --go_out=../go-grpc/hoardremote --go-grpc_out=../go-grpc/hoardremote --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative hoard-remote.proto



