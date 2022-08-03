clean:
	rm -rf  pb/*.go

server:
	go run main.go

gen:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

test:
	go test -cover -race ./...

.PHONY:
	clean server gen test