protoc --go_out=../go-grpc/hoard-data --go-grpc_out=../go-grpc/hoard-data --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative hoard-data.proto
protoc --go_out=../go-grpc/hoard-client --go-grpc_out=../go-grpc/hoard-client --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative hoard-client.proto
protoc --go_out=../go-grpc/hoard-auth --go-grpc_out=../go-grpc/hoard-auth --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative hoard-auth.proto
protoc --go_out=../go-grpc/hoard-compute --go-grpc_out=../go-grpc/hoard-compute --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative hoard-compute.proto
protoc --go_out=../go-grpc/hoard-remote --go-grpc_out=../go-grpc/hoard-remote --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative hoard-remote.proto



