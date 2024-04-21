package repo

import "github.com/DavidG9999/MyProject/domain"

type ProductRepo interface {
	CreateProduct(domain.Product) (*domain.Product, error)
	GetProducts() ([]domain.Product, error)
	UpdateProduct(domain.Product) (*domain.Product, error)
	DeleteProduct(int) error
	ProductById(int) (*domain.Product, error)
}
