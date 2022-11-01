package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	wearablepb "github.com/teo/sirius/proto2"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:5055", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := wearablepb.NewWearableServiceClient(conn)

	stream, err := client.ConsumeBeatsPerMinute(context.Background())
	if err != nil {
		log.Fatalln("Consuming stream", err)
	}

	for i := 0; i < 10; i++ {
		err := stream.Send(&wearablepb.ConsumeBeatsPerMinuteRequest{Uuid: "teo", Value: 100, Minute: uint32(i)})
		if err != nil {
			log.Fatalln("Sending value", err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Closing", err)
	}

	fmt.Println("Total messages", res.GetTotal())
}
