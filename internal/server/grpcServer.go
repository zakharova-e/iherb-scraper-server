package server

import (
	"context"
	"fmt"
	grpcH "github.com/zakharova-e/iherb-scraper-server/internal/catalog/delivery/grpc"
	pb "github.com/zakharova-e/iherb-scraper-server/internal/iherbCatalog"
	"google.golang.org/grpc"
	"log"
	"net"
)

var server = grpc.NewServer()

type IherbGrpcServer struct {
	pb.UnimplementedIHerbCatalogServiceServer
	handler *grpcH.Handler
}

func NewIherbGrpcServer(handler *grpcH.Handler) *IherbGrpcServer {
	if handler == nil {
		panic("grpc handler not specified")
	}
	iherbServer := IherbGrpcServer{}
	pb.RegisterIHerbCatalogServiceServer(server, &iherbServer)
	iherbServer.handler = handler
	return &iherbServer
}

func (s *IherbGrpcServer) Run() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(fmt.Sprintf("failed to listen on port 50051: %v", err))
	}
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}

func (s *IherbGrpcServer) GetProductData(ctx context.Context, request *pb.ProductDataRequest) (*pb.ProductDataResponse, error) {
	return s.handler.GetProductData(ctx, request)
}
