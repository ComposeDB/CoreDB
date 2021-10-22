gen-protoc:
	protoc \
	--go_out=./internal/server/service/ \
	--go-grpc_out=./internal/server/service/ \
	./internal/server/service/protos/*.proto

build:
	go build -o cored ./cli/server/main.go 
	go build -o corectl ./cli/client/main.go