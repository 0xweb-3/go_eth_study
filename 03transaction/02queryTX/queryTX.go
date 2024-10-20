package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

/*
*
用Transactions方法来读取块中的事务，该方法返回一个Transaction类型的列表。 然后，重复遍历集合并获取有关事务的任何信息。
*/
func queryTXByBlock() {
	client, err := ethclient.Dial("https://rpc.payload.de")

	if err != nil {
		fmt.Println(err)
	}

	// 获取块儿信息
	blockNumber := big.NewInt(20998819)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Hash())

	for _, tx := range block.Transactions() {
		// 打印块儿中信息
		fmt.Println(tx.Hash().Hex())
		//fmt.Println(tx.Value().String())
		//fmt.Println(tx.Gas())
		//fmt.Println(tx.GasPrice().Uint64())
		//fmt.Println(tx.Nonce())
		//fmt.Println(tx.Data())
		//fmt.Println(tx.To().Hex())

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		// 发送方的地址
		from, err := types.Sender(types.NewEIP155Signer(chainID), tx)
		if err == nil {
			fmt.Println("发送者地址:", from.Hex())
		}

		// 接收状态
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		//每个事务都有一个收据，其中包含执行事务的结果，例如任何返回值和日志，以及为“1”（成功）或“0”（失败）的事件结果状态。
		fmt.Println(receipt.Status)

		// 打印接收者地址
		if tx.To() != nil {
			fmt.Println("接收者地址:", tx.To().Hex())
		} else {
			fmt.Println("该交易是合约创建交易，没有接收者地址。")
		}
	}
}

/*
*
在不获取块的情况下遍历事务的另一种方法是调用客户端的TransactionInBlock方法。 此方法仅接受块哈希和块内事务的索引值。
*/
func queryTxByHash() {
	client, err := ethclient.Dial("https://rpc.payload.de")

	if err != nil {
		fmt.Println(err)
	}

	blockHash := common.HexToHash("0xcb86ded2de7b87446a5d952feeedb7f14eec3cdb3cc3065646d9ce74ac237f2d") // hash一般使用前都需要做这种包装
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(tx.Hash().Hex()) // 交易的hash

		tx, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex())
		fmt.Println("是否padding状态：", isPending)
	}
}

func main() {
	//queryTXByBlock()

	queryTxByHash()
}
