package balanceChange

import "math/big"

func (bc *BalanceChange) findMaxChangeAddress(addressChanges map[string]*big.Int) string {
	var maxChangeAddress string
	maxChange := new(big.Int)

	for address, change := range addressChanges {
		if change.Abs(change).Cmp(maxChange) > 0 {
			maxChange.Set(change)
			maxChangeAddress = address
		}
	}

	return maxChangeAddress
}
