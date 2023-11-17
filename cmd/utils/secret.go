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

	shares := make([]int64, num_shares)
	var sum int64 = 0

	// GENERATES RANDOM SHARES FOR ALL BUT THE LAST ONE
	for i := int64(0); i < num_shares-1; i++ {
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
func randomShare(max_share int64) (int64, error) {
	if max_share <= 0 {
		return 0, fmt.Errorf("maxShare must be greater than 0")
	}

	max_share_big_int := big.NewInt(int64(max_share))
	share, err := rand.Int(rand.Reader, max_share_big_int)
	if err != nil {
		return 0, err
	}

	return share.Int64(), nil
}

// func SplitSecretMod(secret, num_shares int64) ([]int64, error) {
// 	if num_shares <= 1 {
// 		return nil, fmt.Errorf("number of shares must be greater than 1")
// 	}

// 	if secret < 0 {
// 		return nil, fmt.Errorf("secret must be non-negative")
// 	}

// 	shares := make([]int64, num_shares)

// 	var sum int64 = 0
// 	for i := int64(0); i < num_shares; i++ {
// 		share, err := randomShare(secret - sum)
// 		if err != nil {
// 			return nil, err
// 		}
// 		shares[i] = share
// 		sum += share
// 	}

// 	finalShare := (secret - sum) % num_shares
// 	shares[num_shares-1] = finalShare

// 	return shares, nil
// }
