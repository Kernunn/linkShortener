package main

import (
	"context"
	"fmt"

	//"link_shortener/internal/shortener"
	"github.com/Kernunn/linkShortener/internal/shortener"

	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := shortener.NewLinkShortenerClient(conn)
	for {
		fmt.Print("1 - Create, 2 - Get, 3 - Exit: ")
		var in string
		fmt.Scanln(&in)
		switch in {
		case "1":
			fmt.Print("Input url: ")
			fmt.Scanln(&in)
			resp, err := client.Create(context.Background(), &shortener.LongURL{Url: in})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("short url - %s\n", resp.GetUrl())
		case "2":
			fmt.Print("Input short url: ")
			fmt.Scanln(&in)
			resp, err := client.Get(context.Background(), &shortener.ShortURL{Url: in})
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("url - %s\n", resp.GetUrl())
		case "3":
			return
		default:
			break
		}
	}
}
