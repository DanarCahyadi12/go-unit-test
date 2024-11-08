package test

import (
	"errors"
	"golang-unit-test/entity"
	"golang-unit-test/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProduct(t *testing.T) {
	//struct for each unit test
	type CreateProductTest struct {
		TestName       string                  //testname
		Request        *models.ProductRequest  //request
		ExpectedError  error                   //expected error
		ExpectedResult *models.ProductResponse //expected result
		Mock           func()                  //mocking function
	}

	tests := []CreateProductTest{
		//test number 1
		{
			TestName: "Shouldn't return an error",
			Request: &models.ProductRequest{
				Name:        "product",
				Price:       120,
				Description: "Description",
				Stock:       10,
				Category:    "category",
			},
			ExpectedError: nil,
			ExpectedResult: &models.ProductResponse{
				Name:        "product",
				Price:       120,
				Description: "Description",
				Stock:       10,
				Category:    "category",
			},
			Mock: func() {
				productMock := new(entity.Product)
				productMock.Name = "product"
				productMock.Price = 120
				productMock.Description = "Description"
				productMock.Stock = 10
				productMock.Category = "category"
				productRepositoryMock.Mock.On("CreateOne", productMock).Return(nil) // not return an error
			},
		},

		//test number 2
		{
			TestName: "Should return an error",
			Request: &models.ProductRequest{
				Name:        "product",
				Price:       120,
				Description: "Description",
				Stock:       10,
				Category:    "category",
			},
			ExpectedError:  errors.New("Duplicate product id"),
			ExpectedResult: nil,
			Mock: func() {
				productMock := new(entity.Product)
				productMock.Name = "product"
				productMock.Price = 120
				productMock.Description = "Description"
				productMock.Stock = 10
				productMock.Category = "category"
				productRepositoryMock.Mock.On("CreateOne", productMock).Return(errors.New("Duplicate product id")) // assumming the product id is duplicate
			},
		},

		//Add another test here

	}

	//loop tests
	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			productRepositoryMock.Mock = mock.Mock{} //refresh the mock
			test.Mock()                              //mock the function
			result, err := productServices.CreateProduct(test.Request)
			if test.ExpectedResult != nil { // if the services do not return an error. It means we expect the result
				assert.NotNil(t, result.Id)
				assert.Equal(t, test.ExpectedResult.Name, result.Name)
				assert.Equal(t, test.ExpectedResult.Description, result.Description)
				assert.Equal(t, test.ExpectedResult.Category, result.Category)
				assert.Equal(t, test.ExpectedResult.Price, result.Price)
				assert.Equal(t, test.ExpectedResult.Stock, result.Stock)
				assert.Nil(t, err)
			} else { //it means we expect the error
				assert.Nil(t, result)
				assert.Equal(t, test.ExpectedError.Error(), err.Error())
			}
		})
	}
}
