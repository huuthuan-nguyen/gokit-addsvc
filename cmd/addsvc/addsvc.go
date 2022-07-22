package main

import "flag"

func main() {
	fs := flag.NewFlagSet("addsvc", flag.ExitOnError)
	var (
		debugAddr   = fs.String("debug.addr", ":8080", "Debug and metrics listen address")
		httpAddr    = fs.String("http-addr", ":8081", "HTTP listen address")
		grpcAddr    = fs.String("grpc-addr", ":8082", "gRPC listen address")
		thriftAddr  = fs.String("thrift-addr", ":8083", "Thrift listen address")
		jsonRPCAddr = fs.String("jsonrpc-addr", ":8084", "JSON RPC listen address")
	)
}
