package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// 测试网rpc列表： https://chainlist.org/chain/11155111 或者 https://rpc.info/ethereum-sepolia

/*
*
privateKey 17a01d2d0862c190dd3d286f5233039938c0522da31fd7d580569cdc07e642f4
publicKey a2e624a9cba7cb4b6b2814c8535b11a130f4818027f125c2ecdf23da303491bb822726e209563a1046b20d7dccfebd3857d4b5e9e3d079dcfbf4c9737fc06d18
address 0xEB80a127b2b763C631D8ADCeBb0976b190C8C227
*/
func main() {
	client, err := ethclient.Dial("https://sepolia.drpc.org") // Sepolia 测试网

	if err != nil {
		log.Fatal(err)
	}

	// 加载您的私钥
	privateKeyStr := "17a01d2d0862c190dd3d286f5233039938c0522da31fd7d580569cdc07e642f4"
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获得帐户nonce,回你应该使用的下一个nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(nonce)

	value := big.NewInt(1000000000000000) // 0.001 使用wei
	gasLimit := uint64(21000)             // 一般eth转账 可以这样

	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	From:  fromAddress,
	//	To:    &toAddress,
	//	Value: value,
	//	Data:  data,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 0.0533
	toAddress := common.HexToAddress("0x8ff44C9b5Eab5E5CE8d1d642184b70e9b9587F74")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 生成签名后的交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

}
