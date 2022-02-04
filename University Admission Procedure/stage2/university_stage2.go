package main

/*
University Admission Procedure - Stage 2/7: Raising the bar
https://hyperskill.org/projects/163/stages/844/implement
-------------------------------------------------------------------------------
[Control statements](https://hyperskill.org/learn/topic/1728)
*/

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
