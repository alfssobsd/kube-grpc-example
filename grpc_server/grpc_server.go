package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)
import pb "github.com/alfssobsd/kube-grpc-example/baseproto"

const (
	port = ":9000"
)

type server struct {
	pb.UnimplementedGeneralServiceServer
}

func (s *server) GetServerName(ctx context.Context, in *pb.ServerNameRequest) (*pb.ServerNameResponse, error) {

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	log.Printf("GRPC Server: request from %v", in.ClientName)
	return &pb.ServerNameResponse{Name: name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("GRPC Server: failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGeneralServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("GRPC Server: failed to serve: %v", err)
	}
}
