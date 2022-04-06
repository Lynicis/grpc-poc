package main

import (
	"google.golang.org/grpc"
	"testing"

	"github.com/stretchr/testify/assert"
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

	go func() {
		err := serverInstance.Start()
		if err != nil {
			t.Fail()
		}
	}()
}
