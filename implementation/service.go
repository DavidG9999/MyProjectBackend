package implementation

import (
	"github.com/DavidG9999/MyProject/base"
	"github.com/DavidG9999/MyProject/domain"
	"github.com/DavidG9999/MyProject/repo"
	"github.com/DavidG9999/MyProject/service"
)

type productService struct {
	productRepo repo.ProductRepo
	unitRepo    repo.UnitRepo
}

func NewService(productRepo repo.ProductRepo, unitRepo repo.UnitRepo) service.ProductService {
	return &productService{
		productRepo: productRepo,
		unitRepo:    unitRepo,
	}
}

func (s *productService) CreateProduct(name string, price int, unitId int) (*domain.Product, error) {
	product := domain.Product{Name: name, Price: price, UnitId: unitId}
	createdProduct, err := s.productRepo.CreateProduct(product)
	base.LogError(err)
	return createdProduct, err
}

func (s *productService) GetProducts() ([]domain.Product, error) {
	getProducts, err := s.productRepo.GetProducts()
	base.LogError(err)
	return getProducts, err
}

func (s *productService) UpdateProduct(product domain.Product) (*domain.Product, error) {
	updatedProduct, err := s.productRepo.UpdateProduct(product)
	base.LogError(err)
	return updatedProduct, err
}

func (s *productService) DeleteProduct(id int) error {
	err := s.productRepo.DeleteProduct(id)
	base.LogError(err)
	return err
}

func (s *productService) ProductById(id int) (*domain.Product, error) {
	productById, err := s.productRepo.ProductById(id)
	base.LogError(err)
	return productById, err
}

func (s *productService) CreateUnit(name string) (*domain.Unit, error) {
	unit := domain.Unit{Name: name}
	createdUnit, err := s.unitRepo.CreateUnit(unit)
	base.LogError(err)
	return createdUnit, err
}

func (s *productService) GetUnits() ([]domain.Unit, error) {
	getUnits, err := s.unitRepo.GetUnits()
	base.LogError(err)
	return getUnits, err
}

func (s *productService) UpdateUnit(unit domain.Unit) (*domain.Unit, error) {
	updatedUnit, err := s.unitRepo.UpdateUnit(unit)
	base.LogError(err)
	return updatedUnit, err
}

func (s *productService) DeleteUnit(id int) error {
	err := s.unitRepo.DeleteUnit(id)
	base.LogError(err)
	return err
}

func (s *productService) UnitById(id int) (*domain.Unit, error) {
	unitById, err := s.unitRepo.UnitById(id)
	base.LogError(err)
	return unitById, err
}
