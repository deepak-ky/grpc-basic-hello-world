package main

import (
	"context"
	pb "grpc-new/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.DeepakGreeterServer
}

func (s *server) SayHelloToDeepak(ctx context.Context, req *pb.DeepakHelloRequest) (*pb.DeepakHelloResponse, error) {
	return &pb.DeepakHelloResponse{Greeting: "Hello " + req.GetName() + ", Nice to meet you, Deepak from Deepak server this side."}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Print("server startup failed -> ", err.Error())
	}

	s := grpc.NewServer()

	pb.RegisterDeepakGreeterServer(s, &server{})

	log.Print("Server is running on port ", "9999")

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve due to -> ", err)
	}

}
