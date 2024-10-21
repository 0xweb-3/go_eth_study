# go_eth_study
[中文文档位置](https://goethereumbook.org/zh/client-setup/)
[ETH HD钱包](https://github.com/miguelmota/go-ethereum-hdwallet)

## 安装一些工具
第一步是安装 Solidity编译器 (solc).
```shell
brew update
brew tap ethereum/ethereum
brew install solidity
```

我们还得安装一个叫abigen的工具，来从solidity智能合约生成ABI。
假设您已经在计算机上设置了Go，只需运行以下命令即可安装abigen工具。
```shell 
brew tap ethereum/ethereum
brew install ethereum
abigen --help
```

ABI生成及go描述文件
```shell
cd 04contract/01ABI
solc --abi store.sol # 打印

# 生成 ABI 文件
solc --abi -o ./output Store.sol
# 生成字节码文件
solc --bin -o ./output Store.sol

abigen --bin=./output/Store.bin --abi=./output/Store.abi --pkg=store --out=./store/Store.go
```