package network

import (
	"github.com/r-moraru/tema-TSS/block"
	"github.com/r-moraru/tema-TSS/blockchain"
)

type Network interface {
	GetBlock(string) chan block.Block
	GetBlockchain() chan blockchain.Blockchain
	SendBlock(block.Block)
	GetData() chan string
}
