package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/**********")
	if err != nil {
		log.Fatal(err)
	}

	rawTx := "f86b5e843b9aca24825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802ba0b3fcfb6b08a3c597544ad02390efc7e4b8a1cda12c51ba9f9d9bb96573c10823a02d9b0b8c29197aedbf263ddf30ac241f143433af6270b7a140d4c97c3e17e79f"
   //将原始事务十六进制解码为字节格式
	rawTxBytes, err := hex.DecodeString(rawTx)
   //接下来初始化一个新的types.Transaction指针并从go-ethereumrlp包中调用DecodeBytes，将原始事务字节和指针传递给以太坊事务类型。 RLP是以太坊用于序列化和反序列化数据的编码方法。
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}
