package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// SplitSecret generates shares for a given secret using additive secret sharing.
func SplitSecret(secret int32, numShares int32) ([]int32, error) {
	if numShares <= 1 {
		return nil, fmt.Errorf("numShares must be greater than 1")
	}

	if secret < 0 {
		return nil, fmt.Errorf("secret must be non-negative")
	}

	shares := make([]int32, numShares)
	var sum int32 = 0

	// Generate random shares for all but the last one
	for i := int32(0); i < numShares-1; i++ {
		share, err := randomShare(secret - sum)
		if err != nil {
			return nil, err
		}
		shares[i] = share
		sum += share
	}

	// The last share ensures that the sum of all shares equals the secret
	shares[numShares-1] = secret - sum

	return shares, nil
}

// randomShare generates a random share within the specified range.
func randomShare(maxShare int32) (int32, error) {
	if maxShare <= 0 {
		return 0, fmt.Errorf("maxShare must be greater than 0")
	}

	maxShareBigInt := big.NewInt(int64(maxShare))
	share, err := rand.Int(rand.Reader, maxShareBigInt)
	if err != nil {
		return 0, err
	}

	return int32(share.Int64()), nil
}
