package structs

import (
	"fmt"
)

type Inventory struct {
	items []Item
}

var countStatic int = 1000

func (inv *Inventory) AddItem(item Item) {
	inv.items = append(inv.items, item)
}
func (inv *Inventory) DisplayInventoryReport() {
	fmt.Printf("The length of the array is: %d \n\n", len(inv.items))
	for index, item := range inv.items {
		item.DisplayInfo()
		fmt.Printf("This is the index: %d \n\n", index)
	}
}
func (inv *Inventory) EnterOrder(sku int, quantity int, cost float32) {

	for _, data := range inv.items {
		if data.ReturnSKU() == sku {
			data.NewCost(cost)
		}
	}

}
func (inv *Inventory) EnterSale(sku int, quantity int) SaleActions {
	var saleOrder SaleActions
	for _, data := range inv.items {
		if data.ReturnSKU() == sku {
			quantityReceived := data.UpdateQuantity(quantity)
			var allWeHaveLeftMessage string = ""
			if quantityReceived < quantity {
				allWeHaveLeftMessage = fmt.Sprintf("we only have this %d in our inventory system", quantityReceived)
			}
			saleOrder = Sale{sku: data.ReturnSKU(), quanity: quantityReceived, cost: data.ReturnCost(), message: allWeHaveLeftMessage}
		}
	}
	return saleOrder
}
func Receipt(sales []SaleActions) {
	var sumTotal float32
	for _, data := range sales {
		fmt.Printf("%d   %d %.2f  %.2f %s \n", data.ReturnSKU(), data.ReturnQuantity(), data.ReturnCost(), data.TotalCommerce(), data.ReturnMessageOfQuanity())
		sumTotal += data.TotalCommerce()
	}
	fmt.Printf("Total:          %.2f \n", sumTotal)
	var taxed = 0.0825 * sumTotal
	fmt.Printf("Tax:            %.2f \n", taxed)
	var subTotal = sumTotal + taxed
	fmt.Printf("SubTotal:       $%.2f\n", subTotal)
}
