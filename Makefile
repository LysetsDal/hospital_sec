.PHONY: proto

srv: 
	go run cmd/server/server.go
proto:
    protoc -Iproto/ \
		--go_out=. --go_opt=module=github.com/LysetsDal/hospital_sec \
		--go-grpc_out=. --go-grpc_opt=module=github.com/LysetsDal/hospital_sec \
		proto/*.proto
