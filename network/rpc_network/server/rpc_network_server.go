package server

import (
	"context"

	"github.com/adrianbrad/queue"
	"github.com/r-moraru/tema-TSS/block"
	"github.com/r-moraru/tema-TSS/blockchain"
	pb "github.com/r-moraru/tema-TSS/network/rpc_network/proto"
)

type RpcNetwork struct {
	pb.UnimplementedRpcNetworkServer
	blockQueue               queue.Queue[block.Block]
	blockchainRequestPeerIds queue.Queue[string]
	blockchainQueue          queue.Queue[*blockchain.Blockchain]
}

func (r *RpcNetwork) SendBlock(ctx context.Context, req *pb.Block) (*pb.SendBlockResponse, error) {
	res := &pb.SendBlockResponse{Accepted: false}

	newBlock := block.Block{
		Data:          req.GetData(),
		Previous_hash: req.GetPreviousHash(),
		Timestamp:     req.GetTimestamp(),
		Nonce:         req.GetNonce(),
		Hash:          req.GetHash(),
	}

	if newBlock.CheckHash() {
		res.Accepted = true
		r.blockQueue.Offer(newBlock)
	}

	return res, nil
}

func (r *RpcNetwork) GetBlockchain(ctx context.Context, req *pb.GetBlockchainMessage) (*pb.GetBlockchainResponse, error) {
	res := new(pb.GetBlockchainResponse)

	r.blockchainRequestPeerIds.Offer(req.GetPeerId())

	return res, nil
}

func (r *RpcNetwork) SendBlockchain(ctx context.Context, req *pb.Blockchain) (*pb.SendBlockchainResponse, error) {
	res := &pb.SendBlockchainResponse{Accepted: false}

	newBlockchain := &blockchain.Blockchain{}
	for _, protoBlock := range req.GetBlocks() {
		newBlock := block.Block{
			Data:          protoBlock.GetData(),
			Previous_hash: protoBlock.GetPreviousHash(),
			Timestamp:     protoBlock.GetTimestamp(),
			Nonce:         protoBlock.GetNonce(),
			Hash:          protoBlock.GetHash(),
		}
		newBlockchain.AddBlock(newBlock)
	}

	r.blockchainQueue.Offer(newBlockchain)

	return res, nil
}
