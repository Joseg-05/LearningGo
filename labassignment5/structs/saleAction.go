package structs

type SaleActions interface {
	TotalCommerce() float32
	ReturnSKU() int
	ReturnQuantity() int
	ReturnCost() float32
	ReturnMessageOfQuanity() string
}
