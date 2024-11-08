package mocks

import (
	"golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func NewProductRepositoryMock() *ProductRepositoryMock {
	return &ProductRepositoryMock{
		Mock: mock.Mock{},
	}
}

func (m *ProductRepositoryMock) CreateOne(product *entity.Product) error {
	args := m.Mock.Called(product)

	if err := args.Error(0); err != nil { //if the mock repository return an error
		return err
	}

	return nil

}
