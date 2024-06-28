package balanceChange

import (
	"getblock-test/internal/utils/types/eth"
	"math/big"
	"sync"
)

func (bc *BalanceChange) updateAddressChanges(block *eth.EthBlock, addressChanges map[string]*big.Float, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	ethDecimals := big.NewInt(1e18)

	for _, tx := range block.Transactions {
		if isContractTransaction(tx) {
			continue
		}

		value := new(big.Float)
		weiValue := new(big.Int)
		weiValue.SetString(tx.Value[2:], 16)
		value.SetInt(weiValue)
		value.Quo(value, new(big.Float).SetInt(ethDecimals))

		if _, ok := addressChanges[tx.From]; !ok {
			addressChanges[tx.From] = new(big.Float)
		}
		if _, ok := addressChanges[tx.To]; !ok {
			addressChanges[tx.To] = new(big.Float)
		}

		addressChanges[tx.From].Sub(addressChanges[tx.From], value)
		addressChanges[tx.To].Add(addressChanges[tx.To], value)
	}
}

func isContractTransaction(tx *eth.EthTransaction) bool {
	return tx.To == "" || len(tx.To) == 0
}
