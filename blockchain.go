package main

import (
	"fmt"
	"log"
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
	//bc.transactionPool = transactionPool
	//bc.chain = chain
	//return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

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
	fmt.Printf("previousHash       %s\n", b.transactions)
}

func init() {
	log.SetPrefix("Blockchain:")
}

func main() {
	blockchain := NewBlockchain()
	fmt.Println(blockchain)
}
