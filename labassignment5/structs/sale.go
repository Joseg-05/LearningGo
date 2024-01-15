package structs

type Sale struct {
	sku     int
	quanity int
	cost    float32
	message string
}

func (s Sale) TotalCommerce() float32 {

	return float32(s.cost) * float32(s.quanity)
}
func (s Sale) ReturnSKU() int {
	return s.sku
}
func (s Sale) ReturnQuantity() int {
	return s.quanity
}
func (s Sale) ReturnCost() float32 {
	return s.cost
}

func (s Sale) ReturnMessageOfQuanity() string {
	return s.message
}
