package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
    // 将区块号设置为nil将返回最新的余额。
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 32234371757509808604

	blockNumber := big.NewInt(14569761)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 32234371757509808604

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	//以太坊中的数字是使用尽可能小的单位来处理的，因为它们是定点精度，在ETH中它是wei。要读取ETH值，您必须做计算wei/10^18。
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 32.234371757509808605
    //有时您想知道待处理的账户余额是多少，例如，在提交或等待交易确认后。客户端提供了类似BalanceAt的方法，名为PendingBalanceAt，它接收账户地址作为参数。
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042
}
