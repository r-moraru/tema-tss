package rpc_network

import (
	"testing"

	"github.com/r-moraru/tema-TSS/block"
	"github.com/r-moraru/tema-TSS/blockchain"
	"github.com/r-moraru/tema-TSS/network/rpc_network/mocks"
	"github.com/r-moraru/tema-TSS/network/rpc_network/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewNetwork(t *testing.T) {
	id := "peer0"
	peerMap := map[string]proto.RpcNetworkClient{
		"peer1": mocks.NewRpcNetworkClient(t),
		"peer2": mocks.NewRpcNetworkClient(t),
	}

	rpcNetwork, err := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()

	assert.NoError(t, err, "No error expected.")
	assert.Equal(t, id, rpcNetwork.id)
	assert.Equal(t, peerMap, rpcNetwork.peers)
}

func TestGetBlockReturnsBlockingChanWhenNoBlockAvailable(t *testing.T) {
	id := "peer0"
	peerMap := map[string]proto.RpcNetworkClient{}
	rpcNetwork, _ := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()

	var chanIsBlocking bool
	select {
	case <-rpcNetwork.GetBlock("test_hash"):
		chanIsBlocking = false
	default:
		chanIsBlocking = true
	}

	assert.True(t, chanIsBlocking)
}

func TestGetBlockReturnsBlock(t *testing.T) {
	expectedBlock := block.NewBlock("data1", "test_hash", "00")
	id := "peer0"
	peerMap := map[string]proto.RpcNetworkClient{}
	rpcNetwork, _ := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()
	rpcNetwork.rpcServer.BlockQueue.Offer(expectedBlock)

	var chanIsBlocking bool
	var chanBlock block.Block
	select {
	case chanBlock = <-rpcNetwork.GetBlock("test_hash"):
		chanIsBlocking = false
	default:
		chanIsBlocking = true
	}

	assert.False(t, chanIsBlocking)
	assert.Equal(t, expectedBlock, chanBlock)
}

func TestGetBlockchainReturnsEmptyChannelWhenNoBlockchainAvailable(t *testing.T) {
	id := "peer0"
	peerMap := map[string]proto.RpcNetworkClient{}
	rpcNetwork, _ := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()

	var chanIsBlocking bool
	select {
	case <-rpcNetwork.GetBlockchain():
		chanIsBlocking = false
	default:
		chanIsBlocking = true
	}

	assert.True(t, chanIsBlocking)
}

func TestGetBlockchain(t *testing.T) {
	block1 := block.NewBlock("data1", "", "00")
	block2 := block.NewBlock("data2", block1.Hash, "00")
	expectedBlockchain := blockchain.Blockchain{}
	expectedBlockchain.AddBlock(block1)
	expectedBlockchain.AddBlock(block2)
	id := "peer0"
	peerMap := map[string]proto.RpcNetworkClient{}
	rpcNetwork, _ := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()
	rpcNetwork.rpcServer.BlockchainQueue.Offer(&expectedBlockchain)

	var chanIsBlocking bool
	var chanBlockchain blockchain.Blockchain
	select {
	case chanBlockchain = <-rpcNetwork.GetBlockchain():
		chanIsBlocking = false
	default:
		chanIsBlocking = true
	}

	assert.False(t, chanIsBlocking)
	assert.Equal(t, expectedBlockchain, chanBlockchain)
}

func TestSendBlockchainRequestWithOnePeer(t *testing.T) {
	id := "peer0"
	expectedGetBlockchainRequest := proto.GetBlockchainMessage{PeerId: id}
	getBlockchainResponse := proto.GetBlockchainResponse{}
	peerMap := map[string]proto.RpcNetworkClient{
		"peer1": mocks.NewRpcNetworkClient(t),
	}
	rpcNetwork, _ := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()
	peerMap["peer1"].(*mocks.RpcNetworkClient).EXPECT().GetBlockchain(mock.Anything, &expectedGetBlockchainRequest).Return(&getBlockchainResponse, nil)

	rpcNetwork.SendBlockchainRequest()
}

func TestBlockchainRequestReturnsEmptyChannelWhenNoBlockchainRequestAvailable(t *testing.T) {
	id := "peer0"
	peerMap := map[string]proto.RpcNetworkClient{}
	rpcNetwork, _ := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()

	var chanIsBlocking bool
	select {
	case <-rpcNetwork.BlockchainRequest():
		chanIsBlocking = false
	default:
		chanIsBlocking = true
	}

	assert.True(t, chanIsBlocking)
}

func TestBlockchainRequest(t *testing.T) {
	id := "peer0"
	peerMap := map[string]proto.RpcNetworkClient{}
	rpcNetwork, _ := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()
	rpcNetwork.rpcServer.BlockchainRequestPeerIds.Offer("mock_peer_id")

	var chanIsBlocking bool
	select {
	case <-rpcNetwork.BlockchainRequest():
		chanIsBlocking = false
	default:
		chanIsBlocking = true
	}

	assert.False(t, chanIsBlocking)
}

func TestSendBlock(t *testing.T) {
	id := "peer0"
	block1 := block.NewBlock("data1", "", "00")
	protoBlock1 := proto.Block{
		Data:         block1.Data,
		PreviousHash: block1.Previous_hash,
		Hash:         block1.Hash,
		Nonce:        block1.Nonce,
		Timestamp:    block1.Timestamp,
	}
	expectedPeerResponse := proto.SendBlockResponse{
		Accepted: true,
	}
	peer1 := mocks.NewRpcNetworkClient(t)
	peer2 := mocks.NewRpcNetworkClient(t)
	peerMap := map[string]proto.RpcNetworkClient{
		"peer1": peer1,
		"peer2": peer2,
	}
	peer1.EXPECT().SendBlock(mock.Anything, &protoBlock1).Return(&expectedPeerResponse, nil).Once()
	peer2.EXPECT().SendBlock(mock.Anything, &protoBlock1).Return(&expectedPeerResponse, nil).Once()
	rpcNetwork, _ := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()

	rpcNetwork.SendBlock(block1)
}

func TestSendBlockAddsPeersThatDidNotAcceptBlockToBlockchainRequestQueue(t *testing.T) {
	id := "peer0"
	block1 := block.NewBlock("data1", "", "00")
	protoBlock1 := proto.Block{
		Data:         block1.Data,
		PreviousHash: block1.Previous_hash,
		Hash:         block1.Hash,
		Nonce:        block1.Nonce,
		Timestamp:    block1.Timestamp,
	}
	expectedPeer1Response := proto.SendBlockResponse{
		Accepted: true,
	}
	expectedPeer2Response := proto.SendBlockResponse{
		Accepted: false,
	}
	peer1 := mocks.NewRpcNetworkClient(t)
	peer2 := mocks.NewRpcNetworkClient(t)
	peerMap := map[string]proto.RpcNetworkClient{
		"peer1": peer1,
		"peer2": peer2,
	}
	peer1.EXPECT().SendBlock(mock.Anything, &protoBlock1).Return(&expectedPeer1Response, nil).Once()
	peer2.EXPECT().SendBlock(mock.Anything, &protoBlock1).Return(&expectedPeer2Response, nil).Once()
	peer2.EXPECT().SendBlockchain(mock.Anything, &proto.Blockchain{}).Return(&proto.SendBlockchainResponse{}, nil).Once()
	rpcNetwork, _ := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()

	rpcNetwork.SendBlock(block1)
	rpcNetwork.SendBlockchain(blockchain.Blockchain{})
}

func TestSendBlockchainSendsNothingWhenNoRequestsAvailable(t *testing.T) {
	id := "peer0"
	peer1 := mocks.NewRpcNetworkClient(t)
	peer2 := mocks.NewRpcNetworkClient(t)
	peerMap := map[string]proto.RpcNetworkClient{
		"peer1": peer1,
		"peer2": peer2,
	}
	rpcNetwork, _ := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()

	rpcNetwork.SendBlockchain(blockchain.Blockchain{})
}

func TestSendBlockchainSendsNothingWhenPeerIsNotRecognized(t *testing.T) {
	id := "peer0"
	peer1 := mocks.NewRpcNetworkClient(t)
	peer2 := mocks.NewRpcNetworkClient(t)
	peerMap := map[string]proto.RpcNetworkClient{
		"peer1": peer1,
		"peer2": peer2,
	}
	rpcNetwork, _ := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()
	rpcNetwork.rpcServer.BlockchainRequestPeerIds.Offer("invalid_peer")

	rpcNetwork.SendBlockchain(blockchain.Blockchain{})
}

func TestSendBlockchainSensBlockchainToPeerWithRequest(t *testing.T) {
	id := "peer0"
	block1 := block.NewBlock("data1", "", "00")
	block2 := block.NewBlock("data2", block1.Hash, "00")
	blockchain1 := blockchain.Blockchain{}
	blockchain1.AddBlock(block1)
	blockchain1.AddBlock(block2)
	protoBlock1 := proto.Block{
		Data:         block1.Data,
		PreviousHash: block1.Previous_hash,
		Hash:         block1.Hash,
		Nonce:        block1.Nonce,
		Timestamp:    block1.Timestamp,
	}
	protoBlock2 := proto.Block{
		Data:         block2.Data,
		PreviousHash: block2.Previous_hash,
		Hash:         block2.Hash,
		Nonce:        block2.Nonce,
		Timestamp:    block2.Timestamp,
	}
	protoBlockchain := proto.Blockchain{}
	protoBlockchain.Blocks = append(protoBlockchain.Blocks, &protoBlock1)
	protoBlockchain.Blocks = append(protoBlockchain.Blocks, &protoBlock2)
	peer1 := mocks.NewRpcNetworkClient(t)
	peer2 := mocks.NewRpcNetworkClient(t)
	peerMap := map[string]proto.RpcNetworkClient{
		"peer1": peer1,
		"peer2": peer2,
	}
	peer2.EXPECT().SendBlockchain(mock.Anything, &protoBlockchain).Return(&proto.SendBlockchainResponse{}, nil).Once()
	rpcNetwork, _ := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()
	rpcNetwork.rpcServer.BlockchainRequestPeerIds.Offer("peer2")

	rpcNetwork.SendBlockchain(blockchain1)
}

func TestSendDataReturnsBlockingChannelWhenNothingIsAvailable(t *testing.T) {
	id := "peer0"
	peerMap := map[string]proto.RpcNetworkClient{}
	rpcNetwork, _ := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()

	assert.True(t, rpcNetwork.rpcServer.DataQueue.IsEmpty())

	var chanIsBlocking bool
	select {
	case <-rpcNetwork.GetData():
		chanIsBlocking = false
	default:
		chanIsBlocking = true
	}

	assert.True(t, chanIsBlocking)
}

func TestSendData(t *testing.T) {
	id := "peer0"
	peerMap := map[string]proto.RpcNetworkClient{}
	rpcNetwork, _ := NewRpcNetwork(0, id, peerMap)
	defer rpcNetwork.Stop()
	rpcNetwork.rpcServer.DataQueue.Offer("data1")

	var chanIsBlocking bool
	var receivedData string
	select {
	case receivedData = <-rpcNetwork.GetData():
		chanIsBlocking = false
	default:
		chanIsBlocking = true
	}

	assert.False(t, chanIsBlocking)
	assert.Equal(t, "data1", receivedData)
}
