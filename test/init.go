package test

import (
	"golang-unit-test/services"
	"golang-unit-test/test/mocks"
)

var productRepositoryMock *mocks.ProductRepositoryMock
var productServices *services.ProductService

func init() {
	productRepositoryMock = mocks.NewProductRepositoryMock()
	productServices = services.NewProductService(productRepositoryMock)
}
