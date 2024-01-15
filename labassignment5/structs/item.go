package structs

type Item interface {
	DisplayInfo()
	ReturnSKU() int
	NewCost(cost float32)
	UpdateQuantity(quantity int) int
	ReturnCost() float32
}
