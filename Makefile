.PHONY: proto protoc

srv: 
	go run cmd/server/*

c1:
	go run cmd/client/client.go -own-addr=localhost:8081 -cert=certs/alice.crt \
	-key=certs/alice.key

proto:
	protoc -Iproto/ --go_out=. --go_opt=module=github.com/LysetsDal/hospital_sec \
			--go-grpc_out=. --go-grpc_opt=module=github.com/LysetsDal/hospital_sec \
			proto/*.proto