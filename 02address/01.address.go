package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://rpc.payload.de")
	if err != nil {
		fmt.Println(err)
	}

	address := common.HexToAddress("0x8ff44C9b5Eab5E5CE8d1d642184b70e9b9587F74")
	fmt.Println(address.Hex())
	fmt.Println(address.Bytes())
	fmt.Println(client)

	// 查询账户余额

}
