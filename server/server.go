package main

import "google.golang.org/grpc"

type Server interface {
	GetRPCServer() *grpc.Server
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
