gen:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/cms.proto

# gen:
# 	@protoc --go_out=. --go_opt=paths=source_relative \
# 	--plugin=protoc-gen-go=$(go env GOPATH)/bin/protoc-gen-go \
# 	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
# 	api/cms.proto

# gen:
# 	@protoc --go_out=. --go_opt=paths=source_relative \
# 	--plugin=protoc-gen-go=$(which protoc-gen-go) \
# 	--plugin=protoc-gen-go-grpc=$(which protoc-gen-go-grpc) \
# 	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
# 	api/cms.proto
