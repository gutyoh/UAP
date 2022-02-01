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

	if meanScore/3 >= 60.0 {
		fmt.Println("Congratulations, you are accepted!")
	} else {
		fmt.Println("We regret to inform you that we will not be able to offer you admission.")
	}
}
