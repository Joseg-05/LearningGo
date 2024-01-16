package structs

import (
	"fmt"
)

type DVD struct {
	sku         int
	title       string
	director    string
	leadActor   string
	playingTime string
	genre       string
	cost        float32
	quantity    int
}

func (d DVD) DisplayInfo() {
	fmt.Printf("SKU %d (DVD) \n", d.sku)
	if d.quantity < 10 {
		fmt.Printf("Number of hand:               %d         (Inventory is low, place an order) \n", d.quantity)
	} else {
		fmt.Printf("Number of hand:               %d \n", d.quantity)
	}
	fmt.Printf("Cost:                         %.2f \n", d.cost)
	fmt.Printf("Title:                        %s\n", d.title)
	price := d.cost + d.cost*0.40
	fmt.Printf("Price                         %.2f \n", price)
	total := float32(d.quantity) * price
	fmt.Printf("Total $ value                 %.2f \n", total)
	fmt.Printf("Director:                     %s\n", d.director)
	fmt.Printf("Lead Actor:                   %s\n", d.leadActor)
	fmt.Printf("Playing Time:                 %s\n", d.playingTime)
	fmt.Printf("Genre:                        %s\n", d.genre)
	fmt.Println(".........................................................")
}
func (d DVD) ReturnSKU() int {
	return d.sku
}
func (d *DVD) NewCost(cost float32) {
	fmt.Println(d.cost)
	d.cost = cost
	fmt.Println(d.cost)
}
func (d *DVD) UpdateQuantity(quantity int) int {
	if d.quantity != 0 {
		if d.quantity > quantity {
			d.quantity = d.quantity - quantity
			return quantity
		} else {
			whatisLeft := d.quantity
			d.quantity = 0
			return whatisLeft
		}
	}
	return d.quantity
}
func (d DVD) ReturnCost() float32 {
	return d.cost
}
func CreateADVD() DVD {
	var titleInput string
	var directorInput string
	var leadActorInput string
	var playingTimeInput string
	var genreInput string
	var costInput float32
	var quanityInput int

	fmt.Print("Enter name title of the DVD:")
	fmt.Scan(&titleInput)
	fmt.Print("Enter the director of the DVD")
	fmt.Scan(&directorInput)
	fmt.Print("Enter the lead actor of the DVD:")
	fmt.Scan(&leadActorInput)
	fmt.Print("Enter the playing time of the DVD:")
	fmt.Scan(&playingTimeInput)
	fmt.Print("Enter the genre of the DVD:")
	fmt.Scan(&genreInput)
	fmt.Print("Enter the cost of the DVD:")
	fmt.Scan(&costInput)
	fmt.Print("Enter the quanity of the DVD:")
	fmt.Scan(&quanityInput)
	count := countStatic
	countStatic++
	return DVD{sku: count, title: titleInput, director: directorInput, leadActor: leadActorInput, playingTime: playingTimeInput, genre: genreInput, cost: costInput, quantity: quanityInput}
}
