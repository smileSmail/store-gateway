package rpcx

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8973", "server address")
)

func getRpcClient(servicePath string) (client.XClient, error) {
	d, err := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	if err != nil {
		return nil, err
	}
	xClient := client.NewXClient(servicePath, client.Failtry, client.RandomSelect, d, client.DefaultOption)
	return xClient, nil
}

func CallFuncService(servicePath string, funcName string, req interface{}, resp interface{}) error {
	rpcClient, err := getRpcClient(servicePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := rpcClient.Close(); err != nil {
			panic(err)
		}
	}()
	ctx := context.Background()
	return rpcClient.Call(ctx, funcName, req, resp)
}
