package main

import (
	"google.golang.org/grpc"
	linkShortener "github.com/Kernunn/linkShortener/pkg"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := linkShortener.New()
	linkShortener.RegisterLinkShortenerServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
