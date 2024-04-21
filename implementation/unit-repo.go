package implementation

import (
	"database/sql"

	"github.com/DavidG9999/MyProject/domain"
	"github.com/DavidG9999/MyProject/repo"
)

type unitRepository struct {
	db *sql.DB
}

func NewUnitRepo(db *sql.DB) repo.UnitRepo {
	return &unitRepository{
		db: db,
	}
}

func (r *unitRepository) CreateUnit(unit domain.Unit) (*domain.Unit, error) {
	sql := `INSERT INTO units (name) VALUES ($1) RETURNING id`

	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var insertedID int
	if err := result.QueryRow(unit.Name).Scan(&insertedID); err != nil {
		return nil, err
	}
	return r.UnitById(insertedID)
}

func (r *unitRepository) GetUnits() ([]domain.Unit, error) {
	sql := `SELECT * FROM units`

	result, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	units := []domain.Unit{}

	for result.Next() {
		unit := domain.Unit{}
		if err := result.Scan(&unit.Id, &unit.Name); err != nil {
			return nil, err
		}
		units = append(units, unit)
	}
	return units, nil
}

func (r *unitRepository) UpdateUnit(unit domain.Unit) (*domain.Unit, error) {
	sql := `UPDATE units SET name = $1 WHERE id = $2`

	result, err := r.db.Query(sql, unit.Name, unit.Id)
	if err != nil {
		return nil, err
	}

	result.Close()

	return r.UnitById(unit.Id)
}

func (r *unitRepository) DeleteUnit(id int) error {
	sql := `DELETE FROM units WHERE id = $1`

	_, err := r.db.Exec(sql, id)
	return err
}

func (r *unitRepository) UnitById(id int) (*domain.Unit, error) {
	sql := `SELECT * FROM units WHERE units.id = $1`

	unit := domain.Unit{}

	if err := r.db.QueryRow(sql, id).Scan(&unit.Id, &unit.Name); err != nil {
		return nil, err
	}
	return &unit, nil

}
