package main

// Ganache(正式名称为testrpc)是一个用Node.js编写的以太坊实现，用于在本地开发去中心化应用程序时进行测试。
import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client
}
