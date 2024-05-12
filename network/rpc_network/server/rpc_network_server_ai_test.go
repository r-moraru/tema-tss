package server

import (
	"context"
	"fmt"
	"testing"
	"time"

	pb "github.com/r-moraru/tema-TSS/network/rpc_network/proto"
	"google.golang.org/grpc"
)

func TestRpcServerInitialization(t *testing.T) {
	port := 50051
	server, stopServer, err := RunRpcServer(port)
	if err != nil {
		t.Fatalf("Error starting RpcServer: %v", err)
	}
	defer stopServer()

	// Check if the server is running
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Error connecting to RpcServer: %v", err)
	}
	defer conn.Close()
}

func TestSendBlock(t *testing.T) {
	port := 50052
	server, stopServer, err := RunRpcServer(port)
	if err != nil {
		t.Fatalf("Error starting RpcServer: %v", err)
	}
	defer stopServer()

	// Simulate sending a valid block
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Error connecting to RpcServer: %v", err)
	}
	defer conn.Close()

	client := pb.NewRpcNetworkClient(conn)

	// Define a valid block
	validBlock := &pb.Block{
		Data:         "Test Data",
		PreviousHash: "Previous Hash",
		Timestamp:    time.Now().Unix(),
		Nonce:        12345,
		Hash:         "Valid Hash",
	}

	// Send the block to the server
	response, err := client.SendBlock(context.Background(), validBlock)
	if err != nil {
		t.Fatalf("Error sending block: %v", err)
	}

	// Check if the block was accepted by the server
	if !response.Accepted {
		t.Errorf("Expected block to be accepted, got rejected")
	}
}

func TestSendBlockchain(t *testing.T) {
	port := 50053
	server, stopServer, err := RunRpcServer(port)
	if err != nil {
		t.Fatalf("Error starting RpcServer: %v", err)
	}
	defer stopServer()

	// Simulate sending a valid blockchain
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Error connecting to RpcServer: %v", err)
	}
	defer conn.Close()

	client := pb.NewRpcNetworkClient(conn)

	// Define a valid blockchain
	validBlockchain := &pb.Blockchain{
		Blocks: []*pb.Block{
			{
				Data:         "Test Data 1",
				PreviousHash: "Previous Hash 1",
				Timestamp:    time.Now().Unix(),
				Nonce:        12345,
				Hash:         "Valid Hash 1",
			},
			{
				Data:         "Test Data 2",
				PreviousHash: "Previous Hash 2",
				Timestamp:    time.Now().Unix(),
				Nonce:        67890,
				Hash:         "Valid Hash 2",
			},
		},
	}

	// Send the blockchain to the server
	response, err := client.SendBlockchain(context.Background(), validBlockchain)
	if err != nil {
		t.Fatalf("Error sending blockchain: %v", err)
	}

	// Check if the blockchain was accepted by the server
	if !response.Accepted {
		t.Errorf("Expected blockchain to be accepted, got rejected")
	}
}
