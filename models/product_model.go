package models

type ProductRequest struct {
	Name        string
	Description string
	Category    string
	Price       int
	Stock       int
}
type ProductResponse struct {
	Id          int
	Name        string
	Description string
	Category    string
	Price       int
	Stock       int
}
