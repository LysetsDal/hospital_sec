package util

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc/credentials"
)

func LoadServerTLSConfig(cert string, key string) (credentials.TransportCredentials, error) {
	keyPair, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		return nil, fmt.Errorf("error starting tcp listener: %v", err)
	}
	return keyPair, nil
}

func LoadTLSConfig(c_cert string, c_key string) *tls.Config {
	cert, err := tls.LoadX509KeyPair(c_cert, c_key)
	if err != nil {
		fmt.Printf("could'nt load cert: %v\n", err)
		panic(err)
	}

	return &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true, // Don't use in production!
	}
}

func LoadTLSCert(c_cert string, c_key string) credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair(c_cert, c_key)
	if err != nil {
		fmt.Printf("could'nt load cert: %v\n", err)
		panic(err)
	}

	config := tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true, // Don't use in production!
	}
	return credentials.NewTLS(&config)
}
