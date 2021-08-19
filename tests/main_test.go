package main

import (
	"context"

	"link_shortener/internal/shortener"
	//"github.com/Kernunn/linkShortener/shortener"

	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"strings"
	"testing"
)

func init() {
	s := grpc.NewServer()
	srv := shortener.New()
	shortener.RegisterLinkShortenerServer(s, srv)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789"

func containsForbiddenSymbol(url string) bool {
	for _, c := range url {
		if !strings.Contains(alphabet, string(c)) {
			return true
		}
	}
	return false
}

func TestCreate(t *testing.T) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := shortener.NewLinkShortenerClient(conn)
	resp, err := client.Create(context.Background(), &shortener.LongURL{Url: "example.com"})
	if err != nil {
		log.Fatal(err)
	}
	if len(resp.GetUrl()) != 10 {
		t.Errorf("Expected len = 10, get %d", len(resp.GetUrl()))
	}
	if containsForbiddenSymbol(resp.GetUrl()) {
		t.Errorf("Find forbidden symbol %s", resp.GetUrl())
	}
}

func TestCreateMultiplySameURL(t *testing.T) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := shortener.NewLinkShortenerClient(conn)
	resp1, err := client.Create(context.Background(), &shortener.LongURL{Url: "example1.com"})
	if err != nil {
		log.Fatal(err)
	}
	if len(resp1.GetUrl()) != 10 {
		t.Errorf("Expected len = 10, get %d", len(resp1.GetUrl()))
	}
	if containsForbiddenSymbol(resp1.GetUrl()) {
		t.Errorf("Find forbidden symbol %s", resp1.GetUrl())
	}
	resp2, err := client.Create(context.Background(), &shortener.LongURL{Url: "example1.com"})
	if err != nil {
		log.Fatal(err)
	}
	if len(resp2.GetUrl()) != 10 {
		t.Errorf("Expected len = 10, get %d", len(resp2.GetUrl()))
	}
	if containsForbiddenSymbol(resp2.GetUrl()) {
		t.Errorf("Find forbidden symbol %s", resp2.GetUrl())
	}
	if resp1.GetUrl() != resp2.GetUrl() {
		t.Errorf("Expected same short url, get %s %s", resp1.GetUrl(), resp2.GetUrl())
	}
}

func TestCreateMultiply(t *testing.T) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := shortener.NewLinkShortenerClient(conn)
	responses := make(map[string]struct{})
	for i := 0; i < 10000; i++ {
		resp, err := client.Create(context.Background(), &shortener.LongURL{Url: strconv.Itoa(i)})
		if err != nil {
			log.Fatal(err)
		}
		if len(resp.GetUrl()) != 10 {
			t.Errorf("Expected len = 10, get %d", len(resp.GetUrl()))
		}
		if containsForbiddenSymbol(resp.GetUrl()) {
			t.Errorf("Find forbidden symbol %s", resp.GetUrl())
		}
		if _, ok := responses[resp.GetUrl()]; ok {
			t.Errorf("Not unique short link")
		}
		responses[resp.GetUrl()] = struct{}{}
	}
}

func TestGet(t *testing.T) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := shortener.NewLinkShortenerClient(conn)
	responses := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		resp, err := client.Create(context.Background(), &shortener.LongURL{Url: strconv.Itoa(i)})
		if err != nil {
			log.Fatal(err)
		}
		responses[i] = resp.GetUrl()
	}
	for i := 0; i < 10000; i++ {
		resp, err := client.Get(context.Background(), &shortener.ShortURL{Url: responses[i]})
		if err != nil {
			log.Fatal(err)
		}
		if resp.GetUrl() != strconv.Itoa(i) {
			t.Errorf("Expected %d, get %s", i, resp.GetUrl())
		}
	}
}

func TestWrongGet(t *testing.T) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := shortener.NewLinkShortenerClient(conn)
	_, err = client.Get(context.Background(), &shortener.ShortURL{Url: "test"})
	if err == nil {
		t.Errorf("Expected error, get nil")
	}
}
