package service

import "github.com/DavidG9999/MyProject/domain"

type ProductService interface {
	CreateProduct(name string, price int, unitId int) (*domain.Product, error)
	GetProducts() ([]domain.Product, error)
	UpdateProduct(domain.Product) (*domain.Product, error)
	DeleteProduct(int) error
	ProductById(int) (*domain.Product, error)
	CreateUnit(name string) (*domain.Unit, error)
	GetUnits() ([]domain.Unit, error)
	UpdateUnit(domain.Unit) (*domain.Unit, error)
	DeleteUnit(int) error
	UnitById(int) (*domain.Unit, error)
}
