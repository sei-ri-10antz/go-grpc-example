package server

import (
	"context"
	"io"
	"log"

	"github.com/sei-ri/go-grpc-example/api/proto"
)

type StreamServiceServer struct{}

func (s *StreamServiceServer) Get(ctx context.Context, req *proto.StreamRequest) (*proto.StreamResponse, error) {
	if req.Pt.Value < 0 {
		panic("raise manual panic")
	}
	return &proto.StreamResponse{
		Pt: req.Pt,
	}, nil
}

func (s *StreamServiceServer) List(request *proto.StreamRequest, stream proto.StreamService_ListServer) error {
	for n := 0; n < 10; n++ {
		err := stream.Send(&proto.StreamResponse{
			Pt: &proto.StreamPoint{
				Name:  request.Pt.Name,
				Value: request.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StreamServiceServer) Record(stream proto.StreamService_RecordServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.StreamResponse{Pt: &proto.StreamPoint{Name: "gRPC Stream Server: Record", Value: 1}})
		}
		if err != nil {
			return err
		}

		log.Printf("stream.Recv name: %s, value: %d", req.Pt.Name, req.Pt.Value)
	}

	return nil
}

func (s *StreamServiceServer) Route(stream proto.StreamService_RouteServer) error {
	n := 0
	for {
		err := stream.Send(&proto.StreamResponse{
			Pt: &proto.StreamPoint{
				Name:  "gPRC Stream Client: Route",
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++

		log.Printf("stream.Recv name: %s, value: %d", req.Pt.Name, req.Pt.Value)
	}

	return nil
}
