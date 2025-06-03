package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog/delivery/grpc"
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog/repository/iherbApi"
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog/usecase"
	"github.com/zakharova-e/iherb-scraper-server/internal/config"
	"github.com/zakharova-e/iherb-scraper-server/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
	config.LoadConfig()
	catalogRepo := iherbApi.NewIherbApiRepository(nil)
	catalogUsecase := usecase.NewCatalogUsecase(catalogRepo)
	grpcHandler := grpc.NewGrpcHandler(catalogUsecase)
	srv := server.NewIherbGrpcServer(grpcHandler)
	srv.Run()

}
