package balanceChange

import (
	"getblock-test/internal/getblock"
	"log"
	"math/big"
	"time"
)

type BalanceChange struct {
	getBlockClient getblock.GetBlockClient
}

func NewBalanceChange(getBlockService getblock.GetBlockClient) *BalanceChange {
	return &BalanceChange{
		getBlockClient: getBlockService,
	}
}

func (bc *BalanceChange) GetAddressWithLargestBalanceChange() (string, error) {
	startTime := time.Now()

	lastBlockNumber, err := bc.getLastBlockNumber()
	if err != nil {
		return "", err
	}

	addressChanges, err := bc.processBlocks(lastBlockNumber, 100)
	if err != nil {
		return "", err
	}

	filteredChanges := make(map[string]*big.Float)
	for address, change := range addressChanges {
		if change.Sign() != 0 {
			filteredChanges[address] = change
		}
	}

	if len(filteredChanges) == 0 {
		return "", nil
	}

	maxChangeAddress := ""
	maxChange := new(big.Float)
	for address, change := range filteredChanges {
		if change.Abs(change).Cmp(maxChange) > 0 {
			maxChange.Set(change)
			maxChangeAddress = address
		}
	}

	averageTime := time.Since(startTime) / 100
	log.Printf("Average time is %s", averageTime.String())

	return maxChangeAddress, nil
}
