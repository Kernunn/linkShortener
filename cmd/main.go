package main

import (
	"fmt"
	"github.com/Kernunn/linkShortener/shortener"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := shortener.New()
	shortener.RegisterLinkShortenerServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("api start")
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
