package main

import (
	"fmt"

	"google.golang.org/grpc/credentials"
)

func LoadTLSConfig(cert string, key string) (credentials.TransportCredentials, error) {
    keyPair, err := credentials.NewServerTLSFromFile(cert, key)
    if err != nil {
		return nil, fmt.Errorf("error starting tcp listener: %v", err)
    }
    return keyPair, nil
}
