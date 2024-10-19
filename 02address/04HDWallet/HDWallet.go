package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"math/big"
)

type WalletInfo struct {
	Address  string
	PrivKey  string
	PubKey   string
	Mnemonic string
	Account  accounts.Account
}

func MakeWallet(mnemonic string, index int) (WalletInfo, error) {
	// 创建一个钱包实例
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		panic(err)
	}

	// 解析路径
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", index))
	account, err := wallet.Derive(path, false)
	if err != nil {
		panic(err)
	}

	// 获取私钥
	privKey, err := wallet.PrivateKey(account)
	if err != nil {
		panic(err)
	}

	// 获取公钥
	pubKey := privKey.PublicKey

	// 地址转换为普通地址字符串
	address := account.Address.Hex()

	return WalletInfo{
		Address:  address,
		PrivKey:  common.Bytes2Hex(crypto.FromECDSA(privKey)),
		PubKey:   common.Bytes2Hex(crypto.FromECDSAPub(&pubKey)),
		Mnemonic: mnemonic,
		Account:  account,
	}, nil
}

func signingTransaction(mnemonic string, index int) {
	// 创建一个钱包实例
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		panic(err)
	}

	// 解析路径
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", index))
	account, err := wallet.Derive(path, false)
	if err != nil {
		panic(err)
	}

	// 获取私钥
	privKey, err := wallet.PrivateKey(account)
	if err != nil {
		panic(err)
	}

	nonce := uint64(0)
	value := big.NewInt(1000000000000000000)
	toAddress := common.HexToAddress("0x0")
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(21000000000)
	var data []byte

	// 创建未签名交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// 使用私钥对交易进行签名
	chainID := big.NewInt(1) // 1 代表以太坊主网，测试网可以根据需要设置
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privKey)
	if err != nil {
		panic(err)
	}

	//spew.Dump(signedTx)

	// 打印签名后的交易
	fmt.Printf("Signed Transaction: %v\n", signedTx)
}

func main() {
	mnemonic := "guess category verb rebuild amateur excite fire add bench head blue vital race average swallow material brave spoon museum also mirror lake supreme awful"
	//accountInfo, err := MakeWallet(mnemonic, 0)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(accountInfo)

	signingTransaction(mnemonic, 0)
}
