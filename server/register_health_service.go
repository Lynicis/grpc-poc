package main

import (
	"context"

	"grpc-poc/server/proto/health"
)

type healthService struct {
	health.UnimplementedHealthcheckServiceServer
}

func RegisterHealthService(serverInstance Server) {
	rpcServer := serverInstance.GetRPCServer()
	health.RegisterHealthcheckServiceServer(rpcServer, &healthService{})
}

func (s *healthService) Healthcheck(ctx context.Context, request *health.HealthcheckRequest) (*health.HealthcheckResponse, error) {
	return &health.HealthcheckResponse{Status: "OK"}, nil
}
