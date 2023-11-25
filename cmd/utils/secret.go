package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GENERATES SHARES FOR A GIVEN SECRET USING ADDITIVE SECRET SHARING.
func SplitSecret(secret, num_shares int64) ([]int64, error) {
	if num_shares <= 1 {
		return nil, fmt.Errorf("number of shares must be greater than 1")
	}

	if secret < 0 {
		return nil, fmt.Errorf("secret must be non-negative")
	}

	shares := make([]int64, num_shares - 1)
	var sum int64 = 0

	// GENERATES RANDOM SHARES FOR ALL BUT THE LAST ONE
	for i := int64(0); i < num_shares-1; i++ {
		share := randomShare(secret - sum)

		shares[i] = share
		sum += share
	}

	// THE LAST SHARE ENSURES THAT THE SUM EQUALS THE SECRET
	shares[num_shares-1] = secret - sum

	return shares, nil
}

// GENERATES A RANDOM SHARE WITHIN THE MAX INT RANGE
func randomShare(max_share int64) int64 {
	if max_share <= 0 {
		return 0
	}

	max_share_big_int := big.NewInt(int64(max_share))
	share, _ := rand.Int(rand.Reader, max_share_big_int)

	return share.Int64()
}

func SplitSecretMod(secret, num_shares int64) ([]int64, error) {
	if num_shares <= 1 {
		return nil, fmt.Errorf("number of shares must be greater than 1")
	}

	if secret < 0 {
		return nil, fmt.Errorf("secret must be non-negative")
	}

	shares := make([]int64, num_shares - 1)

	var sum int64 = 0
	for i := int64(0); i < num_shares-1; i++ {
		share := randomShare(secret - sum)

		shares[i] = share
		sum += share
	}

	finalShare := (secret - sum) % 41
	shares = append(shares, finalShare)

	return shares, nil
}
