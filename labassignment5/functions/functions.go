package functions

import (
	"LabAssignments/labassignment5/structs"
	"fmt"
	"strings"
)

func DisplayPrompt() {
	fmt.Println("1)To enter a new Item enter 1")
	fmt.Println("2)To enter an order enter 2")
	fmt.Println("3)For Sale enter 3")
	fmt.Println("4)For complete report enter 4")
	fmt.Println("5)To quit enter 5")
	var menuOption int
	fmt.Scan(&menuOption)
	switch {
	case menuOption == 1:
		addNewItem()
	case menuOption == 2:
		makeAnOrder()
	case menuOption == 3:
		sale()
	case menuOption == 4:
		report()
	case menuOption == 5:
		fmt.Println("Have a nice day!!!!")
	}
}

var inv = structs.Inventory{}

func addNewItem() {
	fmt.Println("B) To add a new book enter B")
	fmt.Println("C) To add a new CD enter C")
	fmt.Println("D) To add a new DVD enter D")
	fmt.Println("M) To return to the main menu enter M")
	var addItemoption string
	fmt.Scan(&addItemoption)

	switch {
	case strings.ToLower(addItemoption) == "b":
		newbook := structs.CreateABook()
		inv.AddItem(&newbook)
		addNewItem()
	case strings.ToLower(addItemoption) == "c":
		newCd := structs.CreateACD()
		inv.AddItem(&newCd)
		addNewItem()
	case strings.ToLower(addItemoption) == "d":
		newDvd := structs.CreateADVD()
		inv.AddItem(&newDvd)
		addNewItem()
	case strings.ToLower(addItemoption) == "m":
		DisplayPrompt()
	default:
		fmt.Println("It's after noon")
	}
}
func makeAnOrder() {
	var SKUItem int = 0
	var quantity int = 0
	var cost float32 = 0.0
	var usercontinue string = ""
	for strings.ToLower(usercontinue) != "q" {
		fmt.Print("Enter SKU for new order:")
		fmt.Scan(&SKUItem)
		fmt.Print("Enter quantity received:")
		fmt.Scan(&quantity)
		fmt.Print("Enter cost:")
		fmt.Scan(&cost)
		inv.EnterOrder(SKUItem, quantity, cost)
		fmt.Print("Next order or q to quit:")
		fmt.Scan(&usercontinue)
	}
	DisplayPrompt()
}
func sale() {
	var sku int
	var quantity int
	var usercontinue string = ""
	var sale []structs.SaleActions
	for strings.ToLower(usercontinue) != "o" {
		fmt.Print("SKU:")
		fmt.Scan(&sku)
		fmt.Print("Quantity:")
		fmt.Scan(&quantity)
		sale = append(sale, inv.EnterSale(sku, quantity))
		fmt.Print("Next SKU or O to total:")
		fmt.Scan(&usercontinue)
	}
	structs.Receipt(sale)
	DisplayPrompt()
}
func report() {
	inv.DisplayInventoryReport()
	DisplayPrompt()
}
