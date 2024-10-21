package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	Store "github.com/0xweb-3/go_eth_study/04contract/01ABI/store"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func Deploy() {
	client, err := ethclient.Dial("https://sepolia.drpc.org") // Sepolia 测试网

	if err != nil {
		panic(err)
	}

	// 加载您的私钥
	privateKeyStr := "17a01d2d0862c190dd3d286f5233039938c0522da31fd7d580569cdc07e642f4"
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	// 使用 Chain ID 11155111 (Sepolia 测试网的 Chain ID)
	chainID := big.NewInt(11155111)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	// 从rpc上获取chainID
	//chainID, err := client.NetworkID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	input := "1.0" // 版本号

	address, tx, instance, err := Store.DeployStore(auth, client, input)
	if err != nil {
		panic(err)
	}

	fmt.Println(address.Hex())   // 0x6247649d7ff9128555c8C9582361012CCFC07A49 合约地址
	fmt.Println(tx.Hash().Hex()) // 0x56432f255e552c74a3fad04f60f11552acb3f3e43d546c8ff2cf3c2b90e7148d 部署交易hash

	_ = instance
}

func loadContract() {
	client, err := ethclient.Dial("https://sepolia.drpc.org") // Sepolia 测试网

	if err != nil {
		panic(err)
	}

	address := common.HexToAddress("0x6247649d7ff9128555c8C9582361012CCFC07A49")
	instance, err := Store.NewStore(address, client)
	if err != nil {
		panic(err)
	}

	fmt.Println("contract is loaded")
	_ = instance

	// 查询合约中的数据
	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"
}

func main() {
	loadContract()
}
