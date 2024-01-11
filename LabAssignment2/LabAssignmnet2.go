package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("This program allows you to enter 5 numbers.")
	fmt.Println("This program can display the highest and lowest")
	fmt.Println("This program can also calculate the average of them")
	fmt.Println("Please enter atleast 2 numbers:")

	var numbers [5]int
	var count int = 0
	var userinput int
	var usercontinue string

	for count < 5 {
		if count == 2 {
			fmt.Println("Would you like to continue with inputing numbers: Y or N")
			fmt.Scan(&usercontinue)
			if strings.ToLower(usercontinue) == "n" {
				break
			} else {
				fmt.Printf("Please enter %d: ", (count + 1))
				fmt.Scan(&userinput)
				numbers[count] = userinput
				count++
			}
		} else {
			fmt.Printf("Please enter %d: ", (count + 1))
			fmt.Scan(&userinput)
			numbers[count] = userinput
			count++
		}
	}
	var highest int = highest(&numbers)
	var lowest int = lowest(&numbers, count)
	var average int = average((&numbers))

	fmt.Printf("The highest number:%d \n", highest)
	fmt.Printf("The lowest number:%d \n", lowest)
	fmt.Printf("The Average:%d \n", average)

	fmt.Println("Would you like to add another 5 numbers: Y or N")
	fmt.Scan(&usercontinue)
	if strings.ToLower(usercontinue) == "y" {
		main()
	}

}
func highest(high *[5]int) int {
	var temphigh int = high[0]
	for index, arr := range high {
		// perform an operation
		if arr > temphigh {
			temphigh = high[index]
		}
	}
	return temphigh
}
func lowest(low *[5]int, count int) int {
	var templow int = low[0]
	var conditional int = 0
	for conditional < count {
		// perform an operation
		if low[conditional] < templow {
			templow = low[conditional]
		}
		conditional++
	}
	return templow
}

func average(avg *[5]int) int {
	var tempsum = 0
	for _, arr := range avg {
		// perform an operation
		tempsum += arr
	}
	return tempsum / len(avg)
}
