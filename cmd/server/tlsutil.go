package hospital

import (
    "google.golang.org/grpc/credentials"
)

func loadTLSConfig(cert string, key string) (credentials.TransportCredentials, error) {
    keyPair, err := credentials.NewServerTLSFromFile(cert, key)
    if err != nil {
        return nil, err
    }
    return keyPair, nil
}
