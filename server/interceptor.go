package server

import (
	"context"
	"log"
	"runtime/debug"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var Interceptor = interceptor{}

type interceptor struct{}

func (interceptor) Logging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		t := time.Now()
		log.Println("received: ", info.FullMethod, req)
		resp, err = handler(ctx, req)
		log.Println("reply: ", info.FullMethod, resp, time.Since(t))
		return
	}
}

func (interceptor) Recovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if e := recover(); e != nil {
				debug.PrintStack()
				err = status.Errorf(codes.Internal, "Panic err: %v", e)
				log.Println("Panic err: ", err)
			}
		}()
		return handler(ctx, req)
	}
}
