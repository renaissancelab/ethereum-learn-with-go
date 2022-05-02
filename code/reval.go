package main

import (
	"bufio"
	token "ethereum-development-with-go/code/contracts_erc20"
"fmt"
"github.com/ethereum/go-ethereum/accounts/abi/bind"
"github.com/ethereum/go-ethereum/common"
"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"log"
"math"
"math/big"
	"os"
	"strings"
	"time"
)

func execute(match string, client *ethclient.Client) {

	time.Sleep(time.Duration(10)*time.Microsecond)
	// Lon (LON) Address
	tokenAddress := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(match)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	if bal.Uint64() > 0 {
		log.Println("balance-not-0:", bal.Uint64()) // "balance: 74605500.647409"
	} else {
		//log.Println("balance-0:", bal.Uint64()) // "balance: 74605500.647409"
	}
	return
    /*
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)         // "name: Star Live Coin"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: STAR"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"
    */
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))


	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}

func main() {
	logFile, err := os.OpenFile("./reval.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	filepath := "/Users/****/Workspace/work/code/blockchain/personal/ethereum-development-with-go-book/code/reval.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/**********")
	if err != nil {
		log.Fatal(err)
	}

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	count := 0
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		arr := strings.Split(line, "#")
		address := arr[1]
		execute(address, client)
		count++
		log.Println(count, line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
}

