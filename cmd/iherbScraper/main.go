package main

import (
	"github.com/joho/godotenv"
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog/delivery/grpc"
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog/repository/iherbApi"
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog/usecase"
	"github.com/zakharova-e/iherb-scraper-server/internal/config"
	"github.com/zakharova-e/iherb-scraper-server/internal/server"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
	config.LoadConfig()
	catalogRepo := iherbApi.NewIherbApiRepository(nil,nil)
	catalogUsecase := usecase.NewCatalogUsecase(catalogRepo)
	grpcHandler := grpc.NewGrpcHandler(catalogUsecase)
	srv := server.NewIherbGrpcServer(grpcHandler)
	srv.Run()

}
