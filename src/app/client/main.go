package main

import (
	"app/helloworld"
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	SayHello(c, name)

	SayHelloHello(c, name)

	SayHelloHelloHello(c, name)

	SayHelloHelloHelloHello(c, name)
}

func SayHello(c helloworld.GreeterClient, name string) {
	r, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

func SayHelloHello(c helloworld.GreeterClient, name string) {
	client, err := c.SayHelloHello(context.Background(), &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for {
		reply, err := client.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Printf("error: %v", err)
		}
		log.Printf("Greeting: %s\n", reply.Message)
	}
}

func SayHelloHelloHello(c helloworld.GreeterClient, name string) {
	client, err := c.SayHelloHelloHello(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	client.Send(&helloworld.HelloRequest{Name: "Hello"})
	client.Send(&helloworld.HelloRequest{Name: "World"})
	client.Send(&helloworld.HelloRequest{Name: name})

	r, err := client.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s\n", r.Message)
}

func SayHelloHelloHelloHello(c helloworld.GreeterClient, name string) {
	client, err := c.SayHelloHelloHelloHello(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	client.Send(&helloworld.HelloRequest{Name: "hatajoe"})
	r, err := client.Recv()
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s\n", r.Message)

	client.Send(&helloworld.HelloRequest{Name: "hatanaka"})
	r, err = client.Recv()
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s\n", r.Message)

	client.Send(&helloworld.HelloRequest{Name: "yusuke"})
	r, err = client.Recv()
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s\n", r.Message)

	if err := client.CloseSend(); err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}
