package main

import (
	math "./utils"
	"fmt"
)

func calculateImportantData() int {
	totalValue := math.Add(3, 5, 7)
	return totalValue
}

func main() {
	fmt.Println("Packages!")
	total := calculateImportantData()
	fmt.Println(total)
}
