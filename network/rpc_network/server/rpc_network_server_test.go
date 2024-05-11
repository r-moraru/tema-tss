package server

import (
	"context"
	"testing"

	"github.com/r-moraru/tema-TSS/block"
	"github.com/r-moraru/tema-TSS/network/rpc_network/proto"
	"github.com/stretchr/testify/assert"
)

func TestServerIsInitializedCorrectly(t *testing.T) {
	s, stopServer, err := RunRpcServer(0)
	defer stopServer()

	assert.NoError(t, err, "Server initialization should not return an error")
	assert.NotNil(t, s.BlockQueue)
	assert.NotNil(t, s.BlockchainQueue)
	assert.NotNil(t, s.DataQueue)
	assert.NotNil(t, s.BlockchainRequestPeerIds)
}

func TestSendBlockAddsBlockToQueue(t *testing.T) {
	b := block.NewBlock("data1", "prev_hash", "00")
	s, stopServer, _ := RunRpcServer(0)
	defer stopServer()

	res, reqErr := s.SendBlock(context.Background(), &proto.Block{
		Data:         b.Data,
		PreviousHash: b.Previous_hash,
		Nonce:        b.Nonce,
		Timestamp:    b.Timestamp,
		Hash:         b.Hash,
	})
	queueBlock, err := s.BlockQueue.Peek()

	assert.NoError(t, reqErr, "No error expected.")
	assert.NoError(t, err, "Block queue should not be empty.")
	assert.Equal(t, b, queueBlock)
	assert.True(t, res.GetAccepted())
}

func TestSendBlockRejectsBlockWithBadHash(t *testing.T) {
	b := block.NewBlock("data1", "prev_hash", "00")
	b.Hash = "bad_hash"
	s, stopServer, _ := RunRpcServer(0)
	defer stopServer()

	res, reqErr := s.SendBlock(context.Background(), &proto.Block{
		Data:         b.Data,
		PreviousHash: b.Previous_hash,
		Nonce:        b.Nonce,
		Timestamp:    b.Timestamp,
		Hash:         b.Hash,
	})
	_, err := s.BlockQueue.Peek()

	assert.NoError(t, reqErr, "No error expected.")
	assert.Error(t, err)
	assert.False(t, res.GetAccepted())
}

func TestGetBlockchainRequest(t *testing.T) {
	peerId := "peerId"
	s, stopServer, _ := RunRpcServer(0)
	defer stopServer()

	_, reqErr := s.GetBlockchain(context.Background(), &proto.GetBlockchainMessage{PeerId: peerId})
	queuePeerId, err := s.BlockchainRequestPeerIds.Peek()

	assert.NoError(t, reqErr, "No error expected.")
	assert.NoError(t, err, "No error expected.")
	assert.Equal(t, peerId, queuePeerId)
}

func TestSendBlockchainAcceptsValidBlockchain(t *testing.T) {
	block1 := block.NewBlock("data1", "", "00")
	block2 := block.NewBlock("data2", block1.Hash, "00")
	sentBlock1 := proto.Block{
		Data:         block1.Data,
		PreviousHash: block1.Previous_hash,
		Nonce:        block1.Nonce,
		Timestamp:    block1.Timestamp,
		Hash:         block1.Hash,
	}
	sentBlock2 := proto.Block{
		Data:         block2.Data,
		PreviousHash: block2.Previous_hash,
		Nonce:        block2.Nonce,
		Timestamp:    block2.Timestamp,
		Hash:         block2.Hash,
	}
	sentBlockchain := &proto.Blockchain{}
	sentBlockchain.Blocks = append(sentBlockchain.Blocks, &sentBlock1)
	sentBlockchain.Blocks = append(sentBlockchain.Blocks, &sentBlock2)
	s, stopServer, _ := RunRpcServer(0)
	defer stopServer()

	res, reqErr := s.SendBlockchain(context.Background(), sentBlockchain)
	queueBlockchain, err := s.BlockchainQueue.Peek()

	assert.NoError(t, reqErr, "No error expected.")
	assert.NoError(t, err, "Queue should not be empty.")
	assert.True(t, res.Accepted)
	assert.Equal(t, 2, queueBlockchain.GetLength())
	queueBlockchainLastBlock, _ := queueBlockchain.GetLastBlock()
	assert.Equal(t, block2, queueBlockchainLastBlock)
}

func TestSendBlockchainIgnoresInvalidBlocks(t *testing.T) {
	block1 := block.NewBlock("data1", "", "00")
	block2 := block.NewBlock("data2", block1.Hash, "00")
	sentBlock1 := proto.Block{
		Data:         block1.Data,
		PreviousHash: block1.Previous_hash,
		Nonce:        block1.Nonce,
		Timestamp:    block1.Timestamp,
		Hash:         block1.Hash,
	}
	sentBlock2 := proto.Block{
		Data:         block2.Data,
		PreviousHash: "bad_previous_hash",
		Nonce:        block2.Nonce,
		Timestamp:    block2.Timestamp,
		Hash:         block2.Hash,
	}
	sentBlockchain := &proto.Blockchain{}
	sentBlockchain.Blocks = append(sentBlockchain.Blocks, &sentBlock1)
	sentBlockchain.Blocks = append(sentBlockchain.Blocks, &sentBlock2)
	s, stopServer, _ := RunRpcServer(0)
	defer stopServer()

	res, reqErr := s.SendBlockchain(context.Background(), sentBlockchain)
	_, err := s.BlockchainQueue.Peek()

	assert.NoError(t, reqErr, "No error expected.")
	assert.Error(t, err, "Queue should be empty.")
	assert.False(t, res.Accepted)
}

func TestSendData(t *testing.T) {
	data := "data1"
	s, stopServer, _ := RunRpcServer(0)
	defer stopServer()

	_, reqErr := s.SendData(context.Background(), &proto.Data{Data: data})
	queueData, err := s.DataQueue.Peek()

	assert.NoError(t, reqErr, "No error expected.")
	assert.NoError(t, err, "Queue should not be empty.")
	assert.Equal(t, data, queueData)
}
