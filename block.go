package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

// Block represents a single block in the blockchain
type Block struct {
	Transaction   string
	Nonce         int
	PreviousHash  string
	Hash          string
}

// Blockchain is a series of connected blocks
var Blockchain []Block

// NewBlock creates a new block with transaction, nonce, and previous hash
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.Hash = CalculateHash(block.Transaction + strconv.Itoa(block.Nonce) + block.PreviousHash)
	Blockchain = append(Blockchain, block)
	return &block
}

// ListBlocks prints all blocks in the blockchain
func ListBlocks() {
	for i, block := range Blockchain {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}

// ChangeBlock changes the transaction of a specific block by index
func ChangeBlock(index int, newTransaction string) {
	if index >= 0 && index < len(Blockchain) {
		Blockchain[index].Transaction = newTransaction
		Blockchain[index].Hash = CalculateHash(Blockchain[index].Transaction + strconv.Itoa(Blockchain[index].Nonce) + Blockchain[index].PreviousHash)
	} else {
		fmt.Println("Invalid block index")
	}
}

// VerifyChain checks the integrity of the blockchain
func VerifyChain() {
	for i := 1; i < len(Blockchain); i++ {
		if Blockchain[i].PreviousHash != Blockchain[i-1].Hash {
			fmt.Println("Blockchain is compromised at block", i)
			return
		}
	}
	fmt.Println("Blockchain is valid")
}

// CalculateHash calculates the hash for a given string
func CalculateHash(stringToHash string) string {
	hash := sha256.New()
	hash.Write([]byte(stringToHash))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}
