.PHONY: proto protoc

srv: 
	go run cmd/server/*

c1:
	go run cmd/client/client.go -port=8080 -name=alice \
	-cert=certs/alice.crt -key=certs/alice.key

c2:
	go run cmd/client/client.go -port=8081 -name=bob \
	-cert=certs/bob.crt -key=certs/bob.key

c3:
	go run cmd/client/client.go -port=8082 -name=charlie \
	-cert=certs/charlie.crt -key=certs/charlie.key

proto:
	protoc -Iproto/ --go_out=. --go_opt=module=github.com/LysetsDal/hospital_sec \
			--go-grpc_out=. --go-grpc_opt=module=github.com/LysetsDal/hospital_sec \
			proto/*.proto