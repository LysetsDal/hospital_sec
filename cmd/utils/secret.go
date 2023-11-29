package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const FIELDSIZE int64 = 100000007


func GetAdditiveShares(secret, N, fieldSize int64) ([]int64, error) {
	// Generate N-1 shares randomly
	shares := make([]int64, N-1)
	for i := int64(0); i < N-1; i++ {
		share, err := randomShare(fieldSize)
		if err != nil {
			return nil, fmt.Errorf("error generating random share: %v", err)
		}
		shares[i] = share
	}

	// Append the final share by subtracting all shares from the secret
	// Modulo is done with fieldSize to ensure the share is within the finite field { 0 .. FIELDSIZE };
	finalShare := (secret - sum(shares)) % fieldSize
	if finalShare < 0 {
		finalShare += fieldSize
	}
	shares = append(shares, finalShare)

	return shares, nil
}

// Generates a random share within the specified field size
func randomShare(fieldSize int64) (int64, error) {
	if fieldSize <= 0 {
		return 0, fmt.Errorf("fieldSize must be greater than 0")
	}

	maxShare := big.NewInt(int64(fieldSize))
	share, err := rand.Int(rand.Reader, maxShare)
	if err != nil {
		fmt.Printf("error generating random number %s\n", err)
	}

	return share.Int64(), nil
}

func sum(values []int64) int64 {
	result := int64(0)
	for _, value := range values {
		result += value
	}
	return result
}

// Helper function to handle modulo for negative numbers
func mod(a, b int64) int64 {
	result := a % b
	if result < 0 {
		result += b
	}
	return result
}

// Hospital can reconstruct secret with this
func ReconstructSecret(data int64) int64 {
	return mod(data, FIELDSIZE)
}
