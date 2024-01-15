package structs

import (
	"fmt"
)

type CD struct {
	sku            int
	title          string
	artistName     string
	numberOfTracks int
	playingTime    string
	genre          string
	cost           float32
	quantity       int
}

func (c CD) DisplayInfo() {
	fmt.Printf("SKU %d (CD) \n", c.sku)
	if c.quantity < 10 {
		fmt.Printf("Number of hand:               %d         (Inventory is low, place an order) \n", c.quantity)
	} else {
		fmt.Printf("Number of hand:               %d \n", c.quantity)
	}
	fmt.Printf("Cost:                         %.2f \n", c.cost)
	fmt.Printf("Title:                        %s\n", c.title)
	price := c.cost + c.cost*0.30
	fmt.Printf("Price                         %.2f \n", price)
	total := float32(c.quantity) * price
	fmt.Printf("Total $ value                 %.2f \n", total)
	fmt.Printf("Artist name:                  %s\n", c.artistName)
	fmt.Printf("Number of Tracks:             %d\n", c.numberOfTracks)
	fmt.Printf("Playing Time:                 %s\n", c.playingTime)
	fmt.Printf("Genre:                        %s\n", c.genre)
	fmt.Println(".........................................................")
}
func (c CD) ReturnSKU() int {
	return c.sku
}
func (c *CD) NewCost(cost float32) {
	c.cost = cost
}
func (c *CD) UpdateQuantity(quantity int) int {
	if c.quantity != 0 {
		if c.quantity > quantity {
			if (quantity - c.quantity) > 0 {
				c.quantity = 0
				return 0
			} else {
				c.quantity = c.quantity - quantity
				return quantity
			}
		}
	}
	return 0
}
func (c CD) ReturnCost() float32 {
	return c.cost
}
func CreateACD() CD {
	var titleInput string
	var artistNameInput string
	var numberOfTrackInput int
	var playingTimeInput string
	var genreInput string
	var costInput float32
	var quanityInput int

	fmt.Print("Enter name title of the CD:")
	fmt.Scan(&titleInput)
	fmt.Print("Enter artist name of the CD:")
	fmt.Scan(&artistNameInput)
	fmt.Print("Enter the number of Tracks in the CD:")
	fmt.Scan(&numberOfTrackInput)
	fmt.Print("Enter the playing time CD:")
	fmt.Scan(&playingTimeInput)
	fmt.Print("Enter the genre of the CD:")
	fmt.Scan(&genreInput)
	fmt.Print("Enter the cost of the CD:")
	fmt.Scan(&costInput)
	fmt.Print("Enter the quantity of this CD:")
	fmt.Scan(&quanityInput)
	var count = countStatic
	countStatic++
	return CD{sku: count, title: titleInput, artistName: artistNameInput, numberOfTracks: numberOfTrackInput, playingTime: playingTimeInput, genre: genreInput, cost: costInput, quantity: quanityInput}
}
