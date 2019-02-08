package rpcconfig

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// RPCServe SERVE
func RPCServe(host string, service interface{}) error {
	rpc.Register(service)
	lis, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	// return nil
}

// RPCClient CLIENT
func RPCClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
