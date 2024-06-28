package main

import (
	"getblock-test/internal/balanceChange"
	"getblock-test/internal/getblock"
	"log"
)

var apiKey = ""

func main() {

	getBlockClient := getblock.NewGetBlockClient(
		apiKey,
	)

	balanceChangeService := balanceChange.NewBalanceChange(*getBlockClient)

	address, err := balanceChangeService.GetAddressWithLargestBalanceChange()
	if err != nil {
		return
	}
	log.Printf("Winner address: %s", address)
}
