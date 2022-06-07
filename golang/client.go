package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "go.protobuf.playground.exchange/play/me/resources/books/v1"
)

func main() {

	// Create a client connection to the books server.
	host := "resources-books-v1-g4upl6muba-ew.a.run.app:443"
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

// NewConn creates a client connection to a gRPC server.
func NewConn(ctx context.Context, host string, insecure bool) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if host != "" {
		opts = append(opts, grpc.WithAuthority(host))
	}

	if insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	return grpc.Dial(host, opts...)
}
