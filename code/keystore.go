package main

// https://goethereumbook.org/zh/keystore/

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs() {
	//go-ethereum中的keystore，每个文件只能包含一个钱包密钥对
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	//调用NewAccount方法创建新的钱包，并给它传入一个用于加密的口令。
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}

func importKs() {
	file := "./tmp/UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	//把文件内容读取成字符串
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
    //password1是原来的加密口令
	password1 := "secret"
	//password2是指定新的加密口令
	password2 := "secret"
	account, err := ks.Import(jsonBytes, password1, password2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
    //导入账户将允许您按期访问该账户，但它将生成新keystore文件！有两个相同的事物是没有意义的，所以我们将删除旧的。
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	createKs()
	//importKs()
}
