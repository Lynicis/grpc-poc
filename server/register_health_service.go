package main

import (
	"context"
	health2 "grpc-poc/proto/health"
)

type healthService struct {
	health2.UnimplementedHealthcheckServiceServer
}

func RegisterHealthService(serverInstance Server) {
	rpcServer := serverInstance.GetRPCServer()
	health2.RegisterHealthcheckServiceServer(rpcServer, &healthService{})
}

func (s *healthService) Healthcheck(ctx context.Context, request *health2.HealthcheckRequest) (*health2.HealthcheckResponse, error) {
	return &health2.HealthcheckResponse{Status: "OK"}, nil
}
