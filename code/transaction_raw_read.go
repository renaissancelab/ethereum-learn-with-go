package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/**********")
	if err != nil {
		log.Fatal(err)
	}

	txHash := common.HexToHash("0xc2688ee9ed1eaac3f71f080b3abe2b5ace80871ef1f9602fba1af0931ea85a98")
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	v, r, s := tx.RawSignatureValues()
	R := r.Bytes()
	S := s.Bytes()
	V := byte(v.Uint64() - 27 + 4)

	sig := make([]byte, 65)
	copy(sig[32-len(R):32], R)
	copy(sig[64-len(S):64], S)
	sig[64] = V
	fmt.Println("V", V)
	_ = V

	//signer := types.HomesteadSigner{}
	signer := types.NewEIP155Signer(tx.ChainId())
	hash := signer.Hash(tx)

	rawTx, err := rlp.EncodeToBytes(tx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", rawTx) //

	publicKeyBytes, err := crypto.Ecrecover(hash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x\n", publicKeyBytes) //

	publicKeyECDSA, err := crypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0xb3de5b46f54d50cfed7fc6af4986d9077fe2ef81
}
