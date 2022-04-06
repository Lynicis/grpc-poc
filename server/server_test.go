package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"grpc-poc/proto/health"
)

func Test_CreateServer(t *testing.T) {
	serverPort := 8090
	serverInstance := CreateServer(serverPort)

	expectedServerStruct := &server{}

	assert.IsType(t, expectedServerStruct, serverInstance)
}

func TestServer_GetRPCServer(t *testing.T) {
	serverPort := 8090
	serverInstance := CreateServer(serverPort)

	getRPCServer := serverInstance.GetRPCServer()

	expectedRPCServer := &grpc.Server{}

	assert.IsType(t, expectedRPCServer, getRPCServer)
}

func TestServer_Start(t *testing.T) {
	serverPort := 8090
	serverInstance := CreateServer(serverPort)

	RegisterHealthService(serverInstance)

	go func() {
		err := serverInstance.Start()
		if err != nil {
			t.Fail()
		}
	}()

	connectionString := fmt.Sprintf(":%d", serverPort)
	connection, err := grpc.Dial(connectionString, grpc.WithInsecure())
	defer connection.Close()

	client := health.NewHealthcheckServiceClient(connection)

	ctx := context.Background()
	response, err := client.Healthcheck(ctx, &health.HealthcheckRequest{})

	assert.Nil(t, err)
	assert.Equal(t, "OK", response.Status)
}
