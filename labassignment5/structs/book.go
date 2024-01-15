package structs

import (
	"fmt"
)

type Book struct {
	sku              int
	title            string
	author           string
	shortDescription string
	cost             float32
	quantity         int
}

// DisplayInfo displays information specific to the Book.
func (b Book) DisplayInfo() {
	// fmt.Sprintf("Book: %s by %s", b.title, b.author)
	fmt.Printf("SKU %d (Book) \n", b.sku)
	if b.quantity < 10 {
		fmt.Printf("Number of hand:               %d         (Inventory is low, place an order) \n", b.quantity)
	} else {
		fmt.Printf("Number of hand:               %d \n", b.quantity)
	}
	fmt.Printf("Cost:                         %.2f \n", b.cost)
	fmt.Printf("Title:                        %s\n", b.title)
	price := b.cost + b.cost*0.25
	fmt.Printf("Price                         %.2f \n", price)
	total := float32(b.quantity) * price
	fmt.Printf("Total $ value                 %.2f \n", total)
	fmt.Printf("Author:                       %s\n", b.author)
	fmt.Printf("Description:                  %s\n", b.shortDescription)
	fmt.Println(".........................................................")
}
func (b Book) ReturnSKU() int {
	return b.sku
}
func (b *Book) NewCost(cost float32) {
	b.cost = cost
}
func (b *Book) UpdateQuantity(quantity int) int {
	if b.quantity != 0 {
		if b.quantity > quantity {
			b.quantity = b.quantity - quantity
			return quantity
		} else {
			whatisLeft := b.quantity
			b.quantity = 0
			return whatisLeft
		}
	}
	return b.quantity
}
func (b Book) ReturnCost() float32 {
	return b.cost
}
func CreateABook() Book {
	var titleInput string
	var authorInput string
	var shortDescriptionInput string
	var costInput float32
	var quanityInput int
	fmt.Print("Enter name title of the book:")
	fmt.Scan(&titleInput)
	fmt.Print("Enter name of the author of the book:")
	fmt.Scan(&authorInput)
	fmt.Print("Enter a short description of the book:")
	fmt.Scan(&shortDescriptionInput)
	fmt.Print("Enter the cost of the book:")
	fmt.Scan(&costInput)
	fmt.Print("Enter the quanity of the book:")
	fmt.Scan(&quanityInput)
	count := countStatic
	countStatic++
	return Book{sku: count, title: titleInput, author: shortDescriptionInput, shortDescription: shortDescriptionInput, cost: costInput, quantity: quanityInput}
}
