run:
	go run main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	$(shell find ./pb -name "*.proto")

clean:
	find pb -name "*.pb.go" -type f -delete

find:
	grpcurl --plaintext -d '{"search": "patia", "status": "deleted"}' localhost:5000 UserService.Find

findone:
	grpcurl --plaintext -d '{"user": "schweinstaiger", "pass": "schweinstaiger", "is_login": true}' localhost:5000 UserService.FindOne