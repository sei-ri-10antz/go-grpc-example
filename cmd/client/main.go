package main

import (
	"context"
	"io"
	"log"

	"github.com/sei-ri/go-grpc-example/api/proto"
	"google.golang.org/grpc"
)

const (
	PORT = "10010"
)

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := proto.NewStreamServiceClient(conn)

	err = callLists(client, &proto.StreamRequest{
		Pt: &proto.StreamPoint{
			Name:  "gRPC Stream Client: list",
			Value: 2020,
		},
	})

	err = callRecord(client, &proto.StreamRequest{
		Pt: &proto.StreamPoint{
			Name:  "gRPC Stream Client: record",
			Value: 2020,
		},
	})

	err = callRoute(client, &proto.StreamRequest{
		Pt: &proto.StreamPoint{
			Name:  "gRPC Stream Client: route",
			Value: 2020,
		},
	})

}

func callLists(client proto.StreamServiceClient, msg *proto.StreamRequest) error {
	stream, err := client.List(context.Background(), msg)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("[response] name: %s, value: %d", resp.Pt.Name, resp.Pt.Value)
	}

	return nil
}

func callRecord(client proto.StreamServiceClient, msg *proto.StreamRequest) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n < 7; n++ {
		err := stream.Send(msg)
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("[response] name: %s, value: %d", resp.Pt.Name, resp.Pt.Value)

	return nil
}

func callRoute(client proto.StreamServiceClient, msg *proto.StreamRequest) error {
	stream, err := client.Route(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n < 5; n++ {
		err = stream.Send(msg)
		if err != nil {
			return err
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("[response] name: %s, value: %d", resp.Pt.Name, resp.Pt.Value)
	}

	stream.CloseSend()

	return nil
}
