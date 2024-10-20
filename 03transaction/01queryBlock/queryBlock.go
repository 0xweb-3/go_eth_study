package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// 获取区块头信息
func getBlockHeader() {
	client, err := ethclient.Dial("https://rpc.payload.de")

	if err != nil {
		fmt.Println(err)
	}

	// 获取最新块头
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String())

	// 获取指定块儿头
	blockNumber := big.NewInt(20998819)
	header, err = client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String())
}

// 获取区块信息
func getBlock() {
	client, err := ethclient.Dial("https://rpc.payload.de")

	if err != nil {
		fmt.Println(err)
	}

	blockNumber := big.NewInt(20998819)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 区块号
	fmt.Println(block.Time())                // 区块时间戳
	fmt.Println(block.Difficulty().Uint64()) // 区块难度
	fmt.Println(block.Hash().Hex())          //
	fmt.Println(len(block.Transactions()))   // 交易列表
}

func main() {
	//getBlockHeader()
	getBlock()
}
