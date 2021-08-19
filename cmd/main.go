package main

import (
	"fmt"

	//"link_shortener/internal/shortener"
	"github.com/Kernunn/linkShortener/internal/shortener"

	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv, err := shortener.New()
	if err != nil {
		log.Fatal(err)
	}
	defer srv.Close()
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
