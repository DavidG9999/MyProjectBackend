package repo

import "github.com/DavidG9999/MyProject/domain"

type UnitRepo interface {
	CreateUnit(domain.Unit) (*domain.Unit, error)
	GetUnits() ([]domain.Unit, error)
	UpdateUnit(domain.Unit) (*domain.Unit, error)
	DeleteUnit(int) error
	UnitById(int) (*domain.Unit, error)
}
