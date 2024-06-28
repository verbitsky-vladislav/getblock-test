package balance_change

import "math/big"

func (bc *BalanceChange) getLastBlockNumber() (*big.Int, error) {
	lastBlockNumberResponse, err := bc.getBlockClient.GetLastBlockNumber()
	if err != nil {
		return nil, err
	}

	lastBlockNumberHex := lastBlockNumberResponse.Result
	lastBlockNumber := new(big.Int)
	lastBlockNumber.SetString(lastBlockNumberHex[2:], 16)

	return lastBlockNumber, nil
}
