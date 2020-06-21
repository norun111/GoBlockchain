package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

type Blockchain struct {
	transactionPool []string
	chain []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

//構造体Blockchainのスライス型のchainにNewBlockメソッドで初期化した構造体を追加している
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

//構造体Blockに値を代入して新しい構造体を初期化
func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp       %d\n", b.timestamp)
	fmt.Printf("nonce       %d\n", b.nonce)
	fmt.Printf("previousHash       %s\n", b.previousHash)
	fmt.Printf("transactions       %s\n", b.transactions)
}

func init() {
	log.SetPrefix("Blockchain:")
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Println("Chain %d \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func main() {
	blockchain := NewBlockchain()
	blockchain.Print()
	blockchain.CreateBlock(5, "hash 1")
	blockchain.Print()
	blockchain.CreateBlock(2, "hash 2")
	blockchain.Print()
}
