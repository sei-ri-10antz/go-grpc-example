package server

import (
	"context"
	"log"
	"net"

	"github.com/sei-ri/go-grpc-example/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Host string `long:"host" description:"Server IP" default:"0.0.0.0" env:"HOST"`
	Port string `long:"port" description:"Server Port" default:"10010" env:"PORT"`
}

func (s *Server) Serve(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.Addr())
	if err != nil {
		return err
	}
	defer lis.Close()

	log.Println("gRPC listening at ", lis.Addr())

	srv := grpc.NewServer()

	proto.RegisterStreamServiceServer(srv, &StreamServiceServer{})
	reflection.Register(srv)

	return srv.Serve(lis)
}

func (s *Server) Addr() string {
	return net.JoinHostPort(s.Host, s.Port)
}
