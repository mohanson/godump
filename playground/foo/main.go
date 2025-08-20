package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Math struct{}

type Args []int
type Result int

func (m *Math) Add(args Args, result *Result) error {
	sum := 0
	for _, v := range args {
		sum += v
	}
	*result = Result(sum)
	return nil
}

type Conn struct {
	io.Reader
	io.Writer
	io.Closer
}

func main() {
	// 注册 RPC 服务
	math := new(Math)
	rpc.Register(math)

	// 设置 HTTP 处理器
	http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		log.Println("DDDDDDDDDDDDDD")
		// conn, _, err := w.(http.Hijacker).Hijack()
		// if err != nil {
		// 	fmt.Println("Hijack error:", err)
		// 	return
		// }
		// defer conn.Close()
		// jsonrpc.ServeConn(conn)

		rwc := Conn{
			Reader: r.Body,
			Writer: w,
			Closer: r.Body,
		}
		jsonrpc.ServeConn(rwc)
		// codec := jsonrpc.NewServerCodec(rwc)
		// rpc.ServeRequest(codec)
	})

	// 启动 HTTP 服务器
	fmt.Println("Server running on 127.0.0.1:8080")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
