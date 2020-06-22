package main

import (
	"io"
	"net/http"
	"strconv"
)

type BlockchainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World")
}

func (bcs *BlockchainServer) Run() {
	server := http.Server{
		Addr: "0.0.0.0:"+strconv.Itoa(int(bcs.port)),
	}
	http.HandleFunc("/", HelloWorld)
	server.ListenAndServe()
}