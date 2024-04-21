package dto

import "github.com/DavidG9999/MyProject/domain"

type (
	CreateUnitRequest struct {
		Name string `json:"name"`
	}
	CreateUnitResponse struct {
		Unit *domain.Unit `json:"unit,omitempty"`
		Err  string       `json:"error,omitempty"`
	}
	GetUnitsRequest struct {
	}
	GetUnitsResponse struct {
		Units []domain.Unit `json:"units,omitempty"`
		Err   string        `json:"error,omitempty"`
	}
	UpdateUnitRequest struct {
		Unit domain.Unit `json:"unit"`
	}
	UpdateUnitResponse struct {
		Unit *domain.Unit `json:"unit,omitempty"`
		Err  string       `json:"error,omitempty"`
	}
	DeleteUnitRequest struct {
		Id int `json:"id"`
	}
	DeleteUnitResponse struct {
		Err string `json:"error,omitempty"`
	}
	UnitByIdRequest struct {
		Id int `json:"id"`
	}
	UnitByIdResponse struct {
		Unit *domain.Unit `json:"unit,omitempty"`
		Err  string       `json:"error,omitempty"`
	}
)
