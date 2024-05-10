package server

import (
	"context"
	"net"
	"strconv"

	"github.com/adrianbrad/queue"
	"github.com/r-moraru/tema-TSS/block"
	"github.com/r-moraru/tema-TSS/blockchain"
	pb "github.com/r-moraru/tema-TSS/network/rpc_network/proto"
	"google.golang.org/grpc"
)

type RpcServer struct {
	pb.UnimplementedRpcNetworkServer
	BlockQueue               queue.Queue[block.Block]
	BlockchainRequestPeerIds queue.Queue[string]
	BlockchainQueue          queue.Queue[*blockchain.Blockchain]
	DataQueue                queue.Queue[string]
}

func (r *RpcServer) SendBlock(ctx context.Context, req *pb.Block) (*pb.SendBlockResponse, error) {
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
		r.BlockQueue.Offer(newBlock)
	}

	return res, nil
}

func (r *RpcServer) GetBlockchain(ctx context.Context, req *pb.GetBlockchainMessage) (*pb.GetBlockchainResponse, error) {
	res := new(pb.GetBlockchainResponse)

	r.BlockchainRequestPeerIds.Offer(req.GetPeerId())

	return res, nil
}

func (r *RpcServer) SendBlockchain(ctx context.Context, req *pb.Blockchain) (*pb.SendBlockchainResponse, error) {
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

	r.BlockchainQueue.Offer(newBlockchain)

	return res, nil
}

func (r *RpcServer) SendData(ctx context.Context, req *pb.Data) (*pb.SendDataResponse, error) {
	r.DataQueue.Offer(req.GetData())
	return &pb.SendDataResponse{}, nil
}

func RunRpcServer(port int) (*RpcServer, func(), error) {
	s := grpc.NewServer()
	rpcServer := RpcServer{
		BlockQueue:               queue.NewLinked([]block.Block{}),
		BlockchainQueue:          queue.NewLinked([]*blockchain.Blockchain{}),
		BlockchainRequestPeerIds: queue.NewLinked([]string{}),
	}
	pb.RegisterRpcNetworkServer(s, &rpcServer)
	stopServerCallback := func() {
		s.Stop()
	}

	lis, err := net.Listen("tcp", "localhost:"+strconv.FormatInt(int64(port), 10))
	if err != nil {
		return &rpcServer, stopServerCallback, err
	}

	go func() {
		err = s.Serve(lis)
	}()

	return &rpcServer, stopServerCallback, err
}
