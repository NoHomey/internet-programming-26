package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"kv/kvproto"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	client := kvproto.NewKVStoreClient(conn)

	getResp, err := client.Get(context.Background(), &kvproto.GetRequest{Key: 9})
	if err != nil {
		panic(err)
	}
	fmt.Println(getResp.GetValue())

	setResp, err := client.Set(context.Background(), &kvproto.SetRequest{Key: 9, Value: 33})
	if err != nil {
		panic(err)
	}
	fmt.Println(setResp.GetValue())

	getResp, err = client.Get(context.Background(), &kvproto.GetRequest{Key: 9})
	if err != nil {
		panic(err)
	}
	fmt.Println(getResp.GetValue())
}
