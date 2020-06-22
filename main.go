package main

import (
	"fmt"
	"goblockchain/wallet"
)

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	w.BlockchainAddress()
}
