package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/mohanson/godump/doa"
)

type Math struct{}

func (m *Math) Add(arg []int, ret *int) error {
	for _, e := range arg {
		*ret += e
	}
	return nil
}

func mainServer() {
	doa.Nil(rpc.Register(&Math{}))
	l := doa.Try(net.Listen("tcp", "127.0.0.1:8080"))
	go func() {
		defer l.Close()
		for {
			c := doa.Try(l.Accept())
			go rpc.ServeCodec(jsonrpc.NewServerCodec(c))
		}
	}()
}

func mainClient() {
	client := doa.Try(jsonrpc.Dial("tcp", "127.0.0.1:8080"))
	defer client.Close()
	ret := 0
	doa.Nil(client.Call("Math.Add", []int{1, 2, 3, 4}, &ret))
	doa.Doa(ret == 10)
}

func main() {
	mainServer()
	mainClient()
}
