package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

// 将地址转换为common.Address类型
func addressToCommonAddress() {
	address := common.HexToAddress("0x8ff44C9b5Eab5E5CE8d1d642184b70e9b9587F74")
	fmt.Println(address.Hex())
	fmt.Println(address.Bytes())
}

// 获取账户余额
func getAddressBalance() {
	// 获取最新的余额
	client, err := ethclient.Dial("https://rpc.payload.de") // ETH主网
	//client, err := ethclient.Dial("wss://arbitrum.callstaticrpc.com") // arb同源链

	account := common.HexToAddress("0x8ff44C9b5Eab5E5CE8d1d642184b70e9b9587F74")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	//指定块儿高时的余额
	blockNumber := big.NewInt(20998819)
	balance, err = client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	// 将wei转换为ETH数量
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	// 待处理的余额
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println("待处理余额", pendingBalance)
}

func main() {
	//addressToCommonAddress()
	getAddressBalance()
}
