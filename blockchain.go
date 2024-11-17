package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Definition of the Block type
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp              %d\n", b.timestamp)
	fmt.Printf("nonce                  %d\n", b.nonce)
	fmt.Printf("previousHash           %s\n", b.previousHash)
	fmt.Printf("transactions           %s\n", b.transactions)
}

// Definition of the type Blockchain
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Sample Previous Hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain: %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 60))
}

func main() {
	blockChain := NewBlockchain()
	blockChain.Print()
	blockChain.CreateBlock(5, "first block's hash")
	blockChain.Print()
	blockChain.CreateBlock(5, "second block's hash")
	blockChain.Print()
}

func init() {
	log.SetPrefix("SundeeepBlockchain ->> ")
}
