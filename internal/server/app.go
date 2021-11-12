package server

import (
	"context"
	"log"
	"net"

	"github.com/composeDB/coredb/internal/schemas"
	pb "github.com/composeDB/coredb/internal/server/service/generated"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedCreateServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) HandleCreate(ctx context.Context, in *pb.CreateRequest) (*pb.CreateReply, error) {
	log.Printf("Received: %v", in.GetName())
	db := schemas.CoreDB{Name: in.GetName()}
	db.Connect()
	return &pb.CreateReply{Message: "Created " + in.GetName()}, nil
}

func Serve() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCreateServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
