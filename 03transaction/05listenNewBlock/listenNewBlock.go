package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	// 注意这里的写法
	client, err := ethclient.Dial("wss://sepolia.drpc.org/ws") // Sepolia 测试网

	headers := make(chan *types.Header)

	// 创建一个新的通道，用于接收最新的区块头
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	//订阅将推送新的区块头事件到我们的通道，因此我们可以使用一个select语句来监听新消息。
	//订阅对象还包括一个error通道，该通道将在订阅失败时发送消息。
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			// 要获得该区块的完整内容，我们可以将区块头的摘要传递给客户端的BlockByHash函数
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}
