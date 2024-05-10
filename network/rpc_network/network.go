package rpc_network

import (
	"context"
	"time"

	"github.com/r-moraru/tema-TSS/block"
	"github.com/r-moraru/tema-TSS/blockchain"
	pb "github.com/r-moraru/tema-TSS/network/rpc_network/proto"
	"github.com/r-moraru/tema-TSS/network/rpc_network/server"
	"google.golang.org/grpc"
)

type RpcNetwork struct {
	rpcServer             *server.RpcServer
	stopRpcServerCallback func()
	id                    string
	peers                 map[string]pb.RpcNetworkClient
}

func NewRpcNetwork(port int, id string, bootstrapPeerAddresses map[string]string) (*RpcNetwork, error) {
	rpcServer, stopRpcServerCallback, err := server.RunRpcServer(port)
	if err != nil {
		return nil, err
	}

	rpcNetwork := &RpcNetwork{
		rpcServer:             rpcServer,
		stopRpcServerCallback: stopRpcServerCallback,
		id:                    id,
		peers:                 make(map[string]pb.RpcNetworkClient),
	}

	for peerId, peerAddress := range bootstrapPeerAddresses {
		conn, err := grpc.Dial(peerAddress)
		if err != nil {
			continue
		}
		defer conn.Close()
		c := pb.NewRpcNetworkClient(conn)
		rpcNetwork.peers[peerId] = c
	}
	return rpcNetwork, nil
}

func (r *RpcNetwork) GetBlock(lastKnownHash string) chan block.Block {
	responseChan := make(chan block.Block, 1)
	if firstBlock, err := r.rpcServer.BlockQueue.Peek(); err == nil {
		responseChan <- firstBlock
	}
	return responseChan
}

func (r *RpcNetwork) GetBlockchain() chan blockchain.Blockchain {
	responseChan := make(chan blockchain.Blockchain, 1)
	if firstBlockchain, err := r.rpcServer.BlockchainQueue.Peek(); err == nil {
		responseChan <- *firstBlockchain
	}
	return responseChan
}

func (r *RpcNetwork) SendBlockchainRequest() {
	for _, peerClient := range r.peers {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		peerClient.GetBlockchain(ctx, &pb.GetBlockchainMessage{PeerId: r.id})
	}
}

func (r *RpcNetwork) BlockchainRequest() chan struct{} {
	responseChan := make(chan struct{}, 1)
	if !r.rpcServer.BlockchainRequestPeerIds.IsEmpty() {
		responseChan <- struct{}{}
	}
	return responseChan
}

func (r *RpcNetwork) SendBlock(b block.Block) {
	for peerId, peerClient := range r.peers {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		sendBlockResponse, err := peerClient.SendBlock(ctx, &pb.Block{
			Data:         b.Data,
			PreviousHash: b.Previous_hash,
			Timestamp:    b.Timestamp,
			Nonce:        b.Nonce,
			Hash:         b.Hash,
		})
		if err != nil {
			continue
		}
		if !sendBlockResponse.Accepted {
			r.rpcServer.BlockchainRequestPeerIds.Offer(peerId)
		}
	}
}

func (r *RpcNetwork) SendBlockchain(b blockchain.Blockchain) {
	peerId, err := r.rpcServer.BlockchainRequestPeerIds.Get()
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	requestBlockchain := pb.Blockchain{}
	for _, b := range b.Blockchain {
		newBlock := pb.Block{
			Data:         b.Data,
			PreviousHash: b.Previous_hash,
			Timestamp:    b.Timestamp,
			Nonce:        b.Nonce,
			Hash:         b.Hash,
		}
		requestBlockchain.Blocks = append(requestBlockchain.Blocks, &newBlock)
	}
	r.peers[peerId].SendBlockchain(ctx, &requestBlockchain)
}

func (r *RpcNetwork) GetData() chan string {
	dataChannel := make(chan string, 1)
	if data, err := r.rpcServer.DataQueue.Get(); err != nil {
		dataChannel <- data
	}
	return dataChannel
}

func (r *RpcNetwork) Stop() {
	r.stopRpcServerCallback()
}
