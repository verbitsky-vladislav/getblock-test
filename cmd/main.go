package main

import (
	balance_change "getblock-test/internal/balance-change"
	"getblock-test/internal/getblock"
	logger "getblock-test/pkg"
)

var apiKey = "3619821a38a74d51b8d70471523eb83f"

func main() {

	getBlockClient := getblock.NewGetBlockClient(
		apiKey,
	)

	balanceChange := balance_change.NewBalanceChange(getBlockClient)

	address, err := balanceChange.GetAddressWithLargestBalanceChange()
	if err != nil {
		return
	}
	logger.Info(address)
}
