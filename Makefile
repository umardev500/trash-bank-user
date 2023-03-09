run:
	go run main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	$(shell find ./pb -name "*.proto")

clean:
	find pb -name "*.pb.go" -type f -delete

find:
	grpcurl --plaintext -d '{"search": "patia", "status": "active"}' localhost:5000 UserService.Find

findone:
	grpcurl --plaintext -d '{"user": "schweinstaiger", "pass": "schweinstaiger", "is_login": true}' localhost:5000 UserService.FindOne

update:
	grpcurl --plaintext -d '{"user_id": "1667875859", "payload": {"user": "jack", "pass": "jack", "details": { "name": "Sarah Schweinstaiger", "email": "sarahsch@gmail.com", "phone": { "number": "083842765573", "is_wa": "yes" }, "address": { "province": "Banten", "city": "Pandeglang", "district": "Patia", "village": "Turus", "postal_code": "42265", "detail": "Kp. Jengkol RT/RW 006/002" }}}}' localhost:5000 UserService.Update