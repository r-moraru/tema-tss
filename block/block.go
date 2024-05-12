package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Data          string
	Previous_hash string
	Timestamp     int64
	Nonce         int64
	Hash          string
}

func NewBlock(data, previous_hash, difficulty string) Block {
	block := Block{
		Data:          data,
		Previous_hash: previous_hash,
		Timestamp:     time.Now().UnixMicro(),
		Nonce:         0,
	}
	block.CalculateHash(difficulty)
	return block
}

func (b *Block) GetHash() string {
	jsonBytes, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", sha256.Sum256(jsonBytes))
}

func (b *Block) CalculateHash(difficulty string) {
	hash := b.GetHash()
	for !strings.HasPrefix(hash, difficulty) {
		b.Nonce++
		hash = b.GetHash()
	}
	b.Hash = hash
}

func (b *Block) CheckHash() bool {
	declaredHash := b.Hash
	b.Hash = ""
	expectedHash := b.GetHash()
	b.Hash = declaredHash
	return declaredHash == expectedHash
}
