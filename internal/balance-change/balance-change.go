package balance_change

import (
	"getblock-test/internal/getblock"
	logger "getblock-test/pkg"
	"math/big"
	"time"
)

type BalanceChange struct {
	getBlockClient getblock.GetBlockClientService
}

func NewBalanceChange(getBlockService getblock.GetBlockClientService) *BalanceChange {
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

	bc.logAddressChanges(filteredChanges)

	averageTime := time.Since(startTime) / 100
	logger.Info("Average execution time per block: ", averageTime)

	return maxChangeAddress, nil
}
