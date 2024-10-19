package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"os"
)

/*
*
keystore是一个包含经过加密了的钱包私钥。go-ethereum中的keystore，每个文件只能包含一个钱包密钥对。
要生成keystore，首先您必须调用NewKeyStore，给它提供保存keystore的目录路径。
然后，您可调用NewAccount方法创建新的钱包，并给它传入一个用于加密的口令。
您每次调用NewAccount，它将在磁盘上生成新的keystore文件。
*/
func createKs(password string) (string, error) {
	ks := keystore.NewKeyStore("./02address/03Keystores/key", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	fmt.Println(account.Address.Hex())
	// 获取生成的 keystore 文件路径
	ksFilePath := account.URL.Path
	return ksFilePath, nil
}

/*
*
现在要导入您的keystore，您基本上像往常一样再次调用NewKeyStore，然后调用Import方法，
该方法接收keystore的JSON数据作为字节。第二个参数是用于加密私钥的口令。
第三个参数是指定一个新的加密口令，但我们在示例中使用一样的口令。导入账户将允许您按期访问该账户，
但它将生成新keystore文件！有两个相同的事物是没有意义的，所以我们将删除旧的。
*/
func importKs(ksFilePath, password string) error {
	ks := keystore.NewKeyStore("./02address/03Keystores/key2", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(ksFilePath)
	if err != nil {
		panic(err)
	}
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		panic(err)
	}

	fmt.Println(account.Address.Hex())

	// 解密并提取私钥
	key, err := keystore.DecryptKey(jsonBytes, password)
	if err != nil {
		panic(err)
	}

	// 打印私钥 (十六进制格式)
	privateKeyHex := fmt.Sprintf("%x", crypto.FromECDSA(key.PrivateKey))
	fmt.Println("Private Key:", privateKeyHex)

	if err := os.Remove(ksFilePath); err != nil {
		panic(err)
	}
	return nil
}

func main() {
	password := "xinBL"
	ksFilePath, err := createKs(password)
	if err != nil {
		panic(err)
	}

	err = importKs(ksFilePath, password)
	if err != nil {
		panic(err)
	}
}
