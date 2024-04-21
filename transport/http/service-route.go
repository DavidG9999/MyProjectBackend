package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/DavidG9999/MyProject/base"
	"github.com/DavidG9999/MyProject/dto"
	"github.com/DavidG9999/MyProject/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewService(svcEndpoints transport.Endpoints, options []kithttp.ServerOption) http.Handler {
	router := mux.NewRouter()

	errorEncoder := kithttp.ServerErrorEncoder(base.EncodeErrorRespornse)

	options = append(options, errorEncoder)

	router.Methods("POST").Path("/products").Handler(
		kithttp.NewServer(
			svcEndpoints.CreateProduct,
			decodeCreateProductRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/products").Handler(
		kithttp.NewServer(
			svcEndpoints.GetProducts,
			decodeGetProductsRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("PUT").Path("/products").Handler(
		kithttp.NewServer(
			svcEndpoints.UpdateProduct,
			decodeUpdateProductRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/products").Handler(
		kithttp.NewServer(
			svcEndpoints.DeleteProduct,
			decodeDeleteProductRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/products/{id}").Handler(
		kithttp.NewServer(
			svcEndpoints.ProductById,
			decodeProductByIdRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("POST").Path("/units").Handler(
		kithttp.NewServer(
			svcEndpoints.CreateUnit,
			decodeCreateUnitRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/units").Handler(
		kithttp.NewServer(
			svcEndpoints.GetUnits,
			decodeGetUnitsRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("PUT").Path("/units").Handler(
		kithttp.NewServer(
			svcEndpoints.UpdateUnit,
			decodeUpdateUnitRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/units").Handler(
		kithttp.NewServer(
			svcEndpoints.DeleteUnit,
			decodeDeleteUnitRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/units/{id}").Handler(
		kithttp.NewServer(
			svcEndpoints.UnitById,
			decodeUnitByIdRequest,
			base.EncodeResponse,
			options...,
		))
	return router
}

func decodeCreateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateProductRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeGetProductsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetProductsRequest
	return req, nil
}

func decodeUpdateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.UpdateProductRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeDeleteProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.DeleteProductRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeProductByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	req := dto.ProductByIdRequest{Id: id}
	return req, err
}

func decodeCreateUnitRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateUnitRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeGetUnitsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetUnitsRequest
	return req, nil
}

func decodeUpdateUnitRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.UpdateUnitRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeDeleteUnitRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.DeleteUnitRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeUnitByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	req := dto.UnitByIdRequest{Id: id}
	return req, err
}
