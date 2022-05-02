package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	txID := common.HexToHash("0x2432ac74f64bbee97fd3cac445e85725cd589524947255b91d6925963077993a")
	receipt, err := client.TransactionReceipt(context.Background(), txID)
	if err != nil {
		log.Fatal(err)
	}

	logID := "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
	for _, vLog := range receipt.Logs {
		fmt.Println(vLog.Topics[0].Hex())
		if vLog.Topics[0].Hex() == logID {
			if len(vLog.Topics) > 2 {
				id := new(big.Int)
				id.SetBytes(vLog.Topics[3].Bytes())

				fmt.Println(id.Uint64()) // 1133
			}
		}
	}
}
