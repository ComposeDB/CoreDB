gen-protoc:
	protoc \
	--go_out=./internal/server/service/ \
	--go-grpc_out=./internal/server/service/ \
	./internal/server/service/protos/*.proto

build:
	go build -o cored ./cli/server/main.go 
	go build -o corectl ./cli/client/main.go

install-deps:
	brew install protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
