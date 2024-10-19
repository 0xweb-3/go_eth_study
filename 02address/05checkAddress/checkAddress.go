package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"regexp"
)

func CheckAddressByRegexp() {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false
}

// 若在该地址存储了字节码，该地址是智能合约。
func checkIsContract(addressStr string) bool {
	client, err := ethclient.Dial("https://rpc.payload.de")
	if err != nil {
		panic(err)
	}
	// 0x Protocol Token (ZRX) smart contract address
	address := common.HexToAddress(addressStr)
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	return isContract
}

func main() {
	CheckAddressByRegexp()
	isContract := checkIsContract("0xe41d2489571d322189246dafa5ebde1f4699f498")
	if isContract { // 不是合约地址就是
		fmt.Println("is contract")
	} else {
		fmt.Println("is account")
	}
}
