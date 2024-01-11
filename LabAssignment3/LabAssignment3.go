package main

import (
	"fmt"
)

func main() {
	fmt.Println("How many contestant are on the contest.")
	var numberOfContestants int
	fmt.Scan(&numberOfContestants)
	display(numberOfContestants)
}
func display(len int) {
	inputNames := make([]string, len)
	inputScores := make([]int, len)
	count := 0
	var nameOfContestant string
	var scoreOfContestant int
	for count < len {
		fmt.Printf("Enter name of contestant %d:", count+1)
		fmt.Scan(&nameOfContestant)
		fmt.Printf("Enter score contestant %d:", count+1)
		fmt.Scan(&scoreOfContestant)
		inputNames[count] = nameOfContestant
		inputScores[count] = scoreOfContestant
		count++
	}
	var q = &inputScores
	var p = &inputNames
	var average float32 = calcAverage(q, len)
	fmt.Printf("%.2f", average)
	printResult(p, q, average, len)
}
func calcAverage(score *[]int, len int) float32 {
	var sum int
	for _, arr := range *score {
		sum += arr
	}
	return float32(float32(sum) / float32(len))
}
func printResult(names *[]string, scores *[]int, average float32, len int) {
	fmt.Println("Names              Score                  Above/Below")
	fmt.Println("__________________________________________________________")
	for index := 0; index < len; index++ {
		name := (*names)[index]
		score := (*scores)[index]
		if float32(score) > average {
			fmt.Printf("%s                   %d                    Above average \n", name, score)
		} else {
			fmt.Printf("%s                   %d                    Below average \n", name, score)
		}
	}
}
