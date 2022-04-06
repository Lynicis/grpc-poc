package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-poc/proto/health"
	"log"
)

func main() {
	connection, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}

	defer func(connection *grpc.ClientConn) {
		err = connection.Close()
		if err != nil {
			log.Panic(err)
		}
	}(connection)

	client := health.NewHealthcheckServiceClient(connection)
	ctx := context.Background()

	response, err := client.Healthcheck(ctx, &health.HealthcheckRequest{})
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(response)
}
