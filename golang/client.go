package main

import (
	"context"
	"log"

	// TODO: ensure your GOPRIVATE golang env is set to point to this private library
	//   GOPRIVATE=external.go.protobuf.playground.exchange
	pb "external.go.protobuf.playground.exchange/play/me/resources/books/v1"
)

func main() {

	// Create a client connection to the books server.
	host := "internal-e0okt3nx.ew.gateway.dev:443"
	conn, err := NewConn(context.Background(), host, false)
	if err != nil {
		log.Fatal(err)
	}

	// Initialise the Books client.
	client := pb.NewBooksServiceClient(conn)

	// Call the ListBooks method on the client library
	books, err := client.ListBooks(context.Background(), &pb.ListBooksRequest{})
	if err != nil {
		log.Println(err)
	}

	log.Println(books)

	// See what other methods are available on the client by typing 'client.'

}
