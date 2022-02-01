package main

import (
	"fmt"
)

func main() {
	var grade float64
	var meanScore float64

	for i := 0; i < 3; i++ {
        fmt.Scanln(&grade)
        meanScore += grade
    }

	fmt.Println(meanScore / 3)
	fmt.Println("Congratulations, you are accepted!")
}
