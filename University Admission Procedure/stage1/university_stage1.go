package main

/*
University Admission Procedure - Stage 1/7: [No one is left behind!](https://hyperskill.org/projects/163/stages/844/implement)
-------------------------------------------------------------------------------
[Introduction to Go](https://hyperskill.org/learn/topic/1829)
[Primitive types](https://hyperskill.org/learn/topic/1807)
[Arithmetic operations](https://hyperskill.org/learn/topic/1785)
[Input/Output](https://hyperskill.org/learn/topic/1506)
[Main (compiling and running)](https://hyperskill.org/learn/topic/1718)
[Loops](https://hyperskill.org/learn/topic/1531)
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
	fmt.Println("Congratulations, you are accepted!")
}
