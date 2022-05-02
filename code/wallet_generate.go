package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	//生成随机私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
    //将其转换为字节, 再次强调一个字节等于2个字符，所以普通的以太坊地址除了0x外是40个字符，20个字节。
	privateKeyBytes := crypto.FromECDSA(privateKey)
	//转换为十六进制字符串，该包提供了一个带有字节切片的Encode方法,然后我们在十六进制编码之后删除“0x”。
	//私钥是256位的二进制字符串，也就是32个字节，64个字符
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	//由于公钥是从私钥派生的,加密私钥具有一个返回公钥的Public方法
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//将其转换为字节
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//将其转换为十六进制的过程与我们使用转化私钥的过程类似。 我们剥离了0x和前2个字符04，它始终是EC前缀，不是必需的。
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	//接受一个ECDSA公钥，并返回公共地址
	//公共地址其实就是公钥的Keccak-256哈希，然后我们取最后40个字符（20个字节）并用“0x”作为前缀
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	//以下是使用 golang.org/x/crypto/sha3 的 Keccak256函数手动完成的方法。
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
}
