package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://sepolia.drpc.org/ws") // Sepolia 测试网

	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress("0x0E8BE547DdE64fB3099b23e127e74E06Ce32c376")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
