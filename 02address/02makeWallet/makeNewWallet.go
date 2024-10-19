package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {
	//生成随机私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	// crypto/ecdsa包使用FromECDSA方法将其转换为字节
	privateKeyBytes := crypto.FromECDSA(privateKey)
	//转换为十六进制字符串,删除“0x”
	fmt.Println("privateKey", hexutil.Encode(privateKeyBytes)[2:]) // 返回的就是私钥

	// 从私钥转换为公钥
	publicKey := privateKey.Public()
	//fmt.Println(publicKey)

	// 将其转换为十六进制，也去除0x
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publicKey", hexutil.Encode(publicKeyBytes)[4:])

	// 生成地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address", address)

	// 公共地址其实就是公钥的Keccak-256哈希，然后我们取最后40个字符（20个字节）并用“0x”作为前缀。
	//hash := sha3.NewLegacyKeccak256()
	//hash.Write(publicKeyBytes[1:])
	//fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}
