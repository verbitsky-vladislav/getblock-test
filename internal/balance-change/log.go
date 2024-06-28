package balance_change

import (
	"fmt"
	logger "getblock-test/pkg"
	"math/big"
)

func (bc *BalanceChange) logAddressChanges(addressChanges map[string]*big.Float) {
	for address, change := range addressChanges {
		sign := "+"
		if change.Sign() < 0 {
			sign = "-"
			change = new(big.Float).Abs(change)
		}
		logger.Info(fmt.Sprintf("Address: %s Balance Change: %s%s", address, sign, change.String()))
	}
}
