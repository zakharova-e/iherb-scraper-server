package grpc

import (
	"context"
	"errors"
	catalog "github.com/zakharova-e/iherb-scraper-server/internal/catalog"
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog/models"
	pb "github.com/zakharova-e/iherb-scraper-server/internal/iherbCatalog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	usecase catalog.ICatalogUsecase
}

func NewGrpcHandler(usecase catalog.ICatalogUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (handler Handler) GetProductData(ctx context.Context, request *pb.ProductDataRequest) (*pb.ProductDataResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request is empty")
	}
	productData, err := handler.usecase.GetProductData(ctx, request.ProductId)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return getProductDataResponse(productData)
}

func getProductDataResponse(productData *models.ProductData) (*pb.ProductDataResponse, error) {
	if productData == nil {
		return nil, errors.New("product data is nil")
	}
	response := pb.ProductDataResponse{
		Name:               productData.DisplayName,
		BasePrice:          productData.ListPriceAmount,
		BasePriceFormatted: productData.ListPrice,
		StockStatus:        getStockStatus(productData.StockStatus),
	}
	return &response, nil
}

func getStockStatus(status int) pb.StockStatusEnum {
	switch status {
	case 0:
		return pb.StockStatusEnum_InStock
	case 1:
		return pb.StockStatusEnum_OutOfStock
	default:
		return pb.StockStatusEnum_UndefinedStatus
	}
}
