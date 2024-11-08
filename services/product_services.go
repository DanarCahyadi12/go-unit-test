package services

import (
	"golang-unit-test/entity"
	"golang-unit-test/models"
	"golang-unit-test/repository"
)

type ProductService struct {
	Repository repository.ProductRepositoryInterface
}

func NewProductService(repository repository.ProductRepositoryInterface) *ProductService {
	return &ProductService{
		Repository: repository,
	}
}

func (s *ProductService) CreateProduct(req *models.ProductRequest) (*models.ProductResponse, error) {
	product := new(entity.Product)
	product.Name = req.Name
	product.Description = req.Description
	product.Category = req.Category
	product.Stock = req.Stock
	product.Price = req.Price

	err := s.Repository.CreateOne(product)
	if err != nil {
		return nil, err
	}

	productResponse := new(models.ProductResponse)
	productResponse.Name = product.Name
	productResponse.Id = product.Id
	productResponse.Description = product.Description
	productResponse.Category = product.Category
	productResponse.Price = product.Price
	productResponse.Stock = product.Stock
	return productResponse, nil
}
