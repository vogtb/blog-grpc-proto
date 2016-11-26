package main

import (
	"log"
	"net"

	pb "../email"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// Used to implement email.EmailServer
type server struct{}

// ReceiveEmail implements email.EmailServer
func (s *server) ReceiveEmail(ctx context.Context, in *pb.Email) (*pb.EmailAck, error) {
	log.Println(in.String())
	return &pb.EmailAck{Status: "OK"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEmailServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
