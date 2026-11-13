package main

import (
	"context"
	"fmt"
	"kv/kvproto"
	"net"

	"google.golang.org/grpc"
)

/*
type KVStoreServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	mustEmbedUnimplementedKVStoreServer()
}
*/

type kvStoreServer struct {
	data map[int64]int64
	kvproto.UnimplementedKVStoreServer
}

func NewKVStoreServer() kvproto.KVStoreServer {
	return &kvStoreServer{data: make(map[int64]int64)}
}

func (server *kvStoreServer) Get(ctx context.Context, in *kvproto.GetRequest) (*kvproto.GetResponse, error) {
	key := in.GetKey()
	val := server.data[key]
	resp := &kvproto.GetResponse{Key: key, Value: val}
	return resp, nil
}

func (server *kvStoreServer) Set(ctx context.Context, in *kvproto.SetRequest) (*kvproto.SetResponse, error) {
	key := in.GetKey()
	val := in.GetValue()
	server.data[key] = val
	resp := &kvproto.SetResponse{Key: key, Value: val}
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()

	kvproto.RegisterKVStoreServer(server, NewKVStoreServer())

	fmt.Println("Preparing to Serve")
	err = server.Serve(listener)
	fmt.Println("Serving")
	if err != nil {
		panic(err)
	}
}
