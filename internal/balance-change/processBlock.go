package balance_change

import (
	logger "getblock-test/pkg"
	"math/big"
	"sync"
)

func (bc *BalanceChange) processBlocks(lastBlockNumber *big.Int, numBlocks int) (map[string]*big.Float, error) {
	addressChanges := make(map[string]*big.Float)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < numBlocks; i++ {
		wg.Add(1)
		go func(blockNumber *big.Int) {
			defer wg.Done()
			blockNumberHex := "0x" + blockNumber.Text(16)
			block, err := bc.getBlockClient.GetBlock(blockNumberHex)
			if err != nil {
				logger.Error(err, "Failed to get block")
				return
			}

			bc.updateAddressChanges(&block.Result, addressChanges, &mu)
		}(new(big.Int).Sub(lastBlockNumber, big.NewInt(int64(i))))
	}

	wg.Wait()

	return addressChanges, nil
}
