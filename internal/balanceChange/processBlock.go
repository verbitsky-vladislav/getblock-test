package balanceChange

import (
	"getblock-test/internal/utils/types"
	"getblock-test/internal/utils/types/eth"
	"math/big"
	"sync"
)

func (bc *BalanceChange) processBlocks(lastBlockNumber *big.Int, numBlocks int) (map[string]*big.Float, error) {
	addressChanges := make(map[string]*big.Float)
	var mu sync.Mutex
	var wg sync.WaitGroup
	const maxRetries = 3
	const maxGoroutines = 100

	sem := make(chan struct{}, maxGoroutines)
	blockChan := make(chan *types.GetBlockBaseResponse[eth.EthBlock], numBlocks)

	// Loop to fetch blocks
	for i := 0; i < numBlocks; i++ {
		wg.Add(1)
		go func(blockNumber *big.Int) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			blockNumberHex := "0x" + blockNumber.Text(16)
			var block *types.GetBlockBaseResponse[eth.EthBlock]
			var err error

			for attempts := 0; attempts < maxRetries; attempts++ {
				block, err = bc.getBlockClient.GetBlock(blockNumberHex)

				if err == nil {
					break
				}
			}

			if err != nil {
				// Handle error fetching block (optional logging or error handling)
				return
			}

			blockChan <- block
		}(new(big.Int).Sub(lastBlockNumber, big.NewInt(int64(i))))
	}

	wg.Wait()
	close(blockChan)

	for block := range blockChan {
		bc.processBlock(&block.Result, addressChanges, &mu)
	}

	return addressChanges, nil
}

func (bc *BalanceChange) processBlock(block *eth.EthBlock, addressChanges map[string]*big.Float, mu *sync.Mutex) {
	bc.updateAddressChanges(block, addressChanges, mu)
}
