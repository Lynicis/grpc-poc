package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type Server interface {
	GetRPCServer() *grpc.Server
	Start() error
}

type server struct {
	port int
	grpc *grpc.Server
}

func CreateServer(port int) Server {
	createRPC := grpc.NewServer()

	return &server{
		port: port,
		grpc: createRPC,
	}
}

func (s *server) GetRPCServer() *grpc.Server {
	return s.grpc
}

func (s *server) Start() error {
	serverPort := fmt.Sprintf(":%d", s.port)
	tcpServer, err := net.Listen("tcp", serverPort)
	if err != nil {
		return err
	}

	return s.grpc.Serve(tcpServer)
}
