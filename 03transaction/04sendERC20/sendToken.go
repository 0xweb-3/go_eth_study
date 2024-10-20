package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
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
const Token = "0x0e8be547dde64fb3099b23e127e74e06ce32c376"

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

	fmt.Println("下一个Nonce", nonce)

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 发送到的地址
	toAddress := common.HexToAddress("0xC3B991ecD6079aCC8493b79aC7691c64Ce09EAC2")
	tokenAddress := common.HexToAddress(Token)

	// 我们现在将从go-ethereum导入crypto/sha3包以生成函数签名的Keccak256哈希。 然后我们只使用前4个字节来获取方法ID。
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	// 我们需要将给我们发送代币的地址左填充到32字节
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	// 确定要发送多少个代币，并且我们需要在big.Int中格式化为wei。
	amount := new(big.Int)
	amount.SetString("100000000000000000000000", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	// 接下来我们只需将方法ID，填充后的地址和填后的转账量，接到将成为我们数据字段的字节片
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 燃气上限制将取决于交易数据的大小和智能合约必须执行的计算步骤。
	//幸运的是，客户端提供了EstimateGas方法，它可以为我们估算所需的燃气量。
	//这个函数从ethereum包中获取CallMsg结构，我们在其中指定数据和地址。 它将返回我们估算的完成交易所需的估计燃气上限。
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gasLimit)

	// 接下来我们需要做的是构建交易事务类型，这类似于您在ETH转账部分中看到的，
	//除了to字段将是代币智能合约地址。 这个常让人困惑。我们还必须在调用中包含0 ETH的值字段和刚刚生成的数据字节。
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit*10, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
