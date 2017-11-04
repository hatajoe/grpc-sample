package main

import (
	"app/helloworld"
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloHello(in *helloworld.HelloRequest, stream helloworld.Greeter_SayHelloHelloServer) error {
	stream.Send(&helloworld.HelloReply{Message: "Hello"})
	stream.Send(&helloworld.HelloReply{Message: "World"})
	stream.Send(&helloworld.HelloReply{Message: in.GetName()})
	return nil
}

func (s *server) SayHelloHelloHello(stream helloworld.Greeter_SayHelloHelloHelloServer) error {
	var buf []string
	for {
		req, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				if err := stream.SendAndClose(&helloworld.HelloReply{Message: strings.Join(buf, ",")}); err != nil {
					log.Printf("error: %v", err)
				}
				break
			}
			log.Printf("error: %v", err)
		}
		buf = append(buf, req.GetName())
	}
	return nil
}

func (s *server) SayHelloHelloHelloHello(stream helloworld.Greeter_SayHelloHelloHelloHelloServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		stream.Send(&helloworld.HelloReply{Message: fmt.Sprintf("Hello %s", req.GetName())})
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
