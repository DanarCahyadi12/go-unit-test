package repository

import (
	"database/sql"
	"golang-unit-test/entity"
)

type ProductRepositoryInterface interface {
	CreateOne(product *entity.Product) error
}

type ProductRepository struct {
	database *sql.DB
}

func NewProductRepository(database *sql.DB) *ProductRepository {
	return &ProductRepository{
		database: database,
	}
}

func (r *ProductRepository) CreateOne(product *entity.Product) error {
	query := `INSERT INTO product VALUES(?,?,?,?,?)`
	_, err := r.database.Exec(query, product.Name, product.Description, product.Category, product.Price, product.Stock)
	if err != nil {
		return err
	}
	return nil
}
