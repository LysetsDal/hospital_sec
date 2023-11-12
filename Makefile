.PHONY: proto protoc

srv: 
	go run cmd/server/*

c1:
	go run cmd/client/client.go -port=6000 -cert=certs/alice.crt \
	-key=certs/alice.key

c2:
	go run cmd/client/client.go -port=7000 -cert=certs/bob.crt \
	-key=certs/bob.key

proto:
	protoc -Iproto/ --go_out=. --go_opt=module=github.com/LysetsDal/hospital_sec \
			--go-grpc_out=. --go-grpc_opt=module=github.com/LysetsDal/hospital_sec \
			proto/*.proto