package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://rpc.payload.de")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(client)
}
