server_greet:
	go run ./greet/server/main.go

client_greet:
	go run ./greet/client/main.go

server_calculator:
	go run ./calculate/server/main.go

client_calculator:
	go run ./calculate/client/main.go


proto_greet:
	protoc --go-grpc_out=./greet/proto --go_out=./greet/proto ./greet/proto/*.proto

proto_calculator:
	protoc --go-grpc_out=./calculate/proto --go_out=./calculate/proto ./calculate/proto/*.proto