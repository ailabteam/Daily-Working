package main
import (
	"bytes"
	"time"
	"strconv"
	"crypto/sha256"
)

// Block keeps block headers
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}

// SetHash calculates and set block hash

func (b* Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), preBlockHash, []byte{}}
	block.SetHash()
	return block
}

// mang byte rong []byte{}
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}