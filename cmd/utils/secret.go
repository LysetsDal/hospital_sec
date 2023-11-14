package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GENERATES SHARES FOR A GIVEN SECRET USING ADDITIVE SECRET SHARING.
func SplitSecret(secret int32, num_shares int32) ([]int32, error) {
	if num_shares <= 1 {
		return nil, fmt.Errorf("number of shares must be greater than 1")
	}

	if secret < 0 {
		return nil, fmt.Errorf("secret must be non-negative")
	}

	shares := make([]int32, num_shares)
	var sum int32 = 0

	// GENERATES RANDOM SHARES FOR ALL BUT THE LAST ONE
	for i := int32(0); i < num_shares-1; i++ {
		share, err := randomShare(secret - sum)
		if err != nil {
			return nil, err
		}
		shares[i] = share
		sum += share
	}

	// THE LAST SHARE ENSURES THAT THE SUM EQUALS THE SECRET
	shares[num_shares-1] = secret - sum

	return shares, nil
}

// GENERATES A RANDOM SHARE WITHIN THE GIVEN RANGE
func randomShare(max_share int32) (int32, error) {
	if max_share <= 0 {
		return 0, fmt.Errorf("maxShare must be greater than 0")
	}

	max_share_big_int := big.NewInt(int64(max_share))
	share, err := rand.Int(rand.Reader, max_share_big_int)
	if err != nil {
		return 0, err
	}

	return int32(share.Int64()), nil
}
