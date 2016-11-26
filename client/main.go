package main

import (
	"fmt"
	"log"
	"time"

	pb "../email"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting client.")
	email := pb.Email{
		To:        "you",
		From:      "me",
		Body:      "Hey there, protobufs are cool!",
		Subject:   "Daily reminder",
		Timestamp: int64(time.Now().Unix() * 1000),
	}
	fmt.Println(email.String())
	out, _ := proto.Marshal(&email)
	fmt.Println(out)
	send(email)
}

func send(email pb.Email) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEmailServiceClient(conn)

	r, err := c.ReceiveEmail(context.Background(), &email)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Status)
}
