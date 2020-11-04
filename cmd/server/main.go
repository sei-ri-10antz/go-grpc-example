package main

import (
	"context"
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/sei-ri/go-grpc-example/server"
)

func main() {
	srv := server.Server{}

	parser := flags.NewParser(&srv, flags.Default)
	parser.ShortDescription = `gRPC-Steraming`
	parser.LongDescription = `Options for gRPC-Steraming`

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	if err := srv.Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
