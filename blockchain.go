package main

import (
	"fmt"
	"log"
	"time"
)

func init() {
	log.SetPrefix("Blockchain:")
}

func main() {
	log.Println("test")
	fmt.Println("test2")
}

type Block struct {
	nonce int
	previousHash string
	timestamp int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}