package transport

import (
	"context"

	"github.com/DavidG9999/MyProject/dto"
	"github.com/DavidG9999/MyProject/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateProduct endpoint.Endpoint
	GetProducts   endpoint.Endpoint
	UpdateProduct endpoint.Endpoint
	DeleteProduct endpoint.Endpoint
	ProductById   endpoint.Endpoint
	CreateUnit    endpoint.Endpoint
	GetUnits      endpoint.Endpoint
	UpdateUnit    endpoint.Endpoint
	DeleteUnit    endpoint.Endpoint
	UnitById      endpoint.Endpoint
}

func MakeEndpoints(s service.ProductService) Endpoints {
	return Endpoints{
		CreateProduct: makeCreateProductEndpoint(s),
		GetProducts:   makeGetProductsEndpoint(s),
		UpdateProduct: makeUpdateProductEndpoint(s),
		DeleteProduct: makeDeleteProductEndpoint(s),
		ProductById:   makeProductByIdEndpoint(s),
		CreateUnit:    makeCreateUnitEndpoint(s),
		GetUnits:      makeGetUnitsEndpoint(s),
		UpdateUnit:    makeUpdateUnitEndpoint(s),
		DeleteUnit:    makeDeleteUnitEndpoint(s),
		UnitById:      makeUnitByIdEndpoint(s),
	}
}

func makeCreateProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CreateProductRequest)

		product, err := s.CreateProduct(req.Name, req.Price, req.UnitId)

		return dto.CreateProductResponse{Product: product}, err
	}
}

func makeGetProductsEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		products, err := s.GetProducts()

		return dto.GetProductsResponse{Products: products}, err
	}
}

func makeUpdateProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UpdateProductRequest)
		product, err := s.UpdateProduct(req.Product)
		return dto.UpdateProductResponse{Product: product}, err
	}
}

func makeDeleteProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.DeleteProductRequest)
		err := s.DeleteProduct(req.Id)
		return dto.DeleteProductResponse{}, err
	}
}

func makeProductByIdEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.ProductByIdRequest)
		product, err := s.ProductById(req.Id)
		return dto.ProductByIdResponse{Product: product}, err
	}
}

func makeCreateUnitEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CreateUnitRequest)
		unit, err := s.CreateUnit(req.Name)
		return dto.CreateUnitResponse{Unit: unit}, err
	}
}

func makeGetUnitsEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		units, err := s.GetUnits()
		return dto.GetUnitsResponse{Units: units}, err
	}
}

func makeUpdateUnitEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UpdateUnitRequest)
		unit, err := s.UpdateUnit(req.Unit)
		return dto.UpdateUnitResponse{Unit: unit}, err
	}
}

func makeDeleteUnitEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.DeleteUnitRequest)
		err := s.DeleteUnit(req.Id)
		return dto.DeleteUnitResponse{}, err
	}
}

func makeUnitByIdEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UnitByIdRequest)
		unit, err := s.UnitById(req.Id)
		return dto.UnitByIdResponse{Unit: unit}, err
	}
}
