package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	rawTx := "f86d82069e843b9aca16825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802ba0cbd7194eeebccae26b92033d7495c2f63940afae88d4ae7e69af4617b3ea8b79a042ab7bed0b6a37b8e550c3a681dcd16a3eba037dc8bddeb91a2b3aff13c96141"

	tx := new(types.Transaction)
	rawTxBytes, err := hex.DecodeString(rawTx)
	rlp.DecodeBytes(rawTxBytes, &tx)

	fmt.Println(tx.Hash().Hex())

	msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg.From().Hex())
	// 0x96216849c49358B10257cb55b28eA603c874b05E
}
