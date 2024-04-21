package implementation

import (
	"database/sql"

	"github.com/DavidG9999/MyProject/domain"
	"github.com/DavidG9999/MyProject/repo"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) repo.ProductRepo {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) CreateProduct(product domain.Product) (*domain.Product, error) {
	sql := `INSERT INTO products (name, price, unit_id) 
	VALUES ($1, $2, $3) RETURNING id`

	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var insertedID int
	if err := result.QueryRow(product.Name, product.Price, product.UnitId).Scan(&insertedID); err != nil {
		return nil, err
	}
	return r.ProductById(insertedID)

}

func (r *productRepository) GetProducts() ([]domain.Product, error) {
	sql := `SELECT * FROM products`

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []domain.Product{}

	for rows.Next() {
		product := domain.Product{}
		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.UnitId); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *productRepository) UpdateProduct(product domain.Product) (*domain.Product, error) {
	sql := `UPDATE products SET name = $1, price = $2, unit_id = $3 WHERE id = $4`

	result, err := r.db.Query(sql, product.Name, product.Price, product.UnitId, product.Id)
	if err != nil {
		return nil, err
	}
	result.Close()
	return r.ProductById(product.Id)

}

func (r *productRepository) DeleteProduct(id int) error {
	sql := `DELETE FROM products WHERE id = $1`

	_, err := r.db.Exec(sql, id)
	return err
}

func (r *productRepository) ProductById(id int) (*domain.Product, error) {
	sql := `SELECT * FROM products WHERE products.id = $1`

	product := domain.Product{}

	if err := r.db.QueryRow(sql, id).Scan(&product.Id, &product.Name, &product.Price, &product.UnitId); err != nil {
		return nil, err
	}
	return &product, nil
}
