package main

/*
University Admission Procedure - Stage 3/7: Going big
https://hyperskill.org/projects/163/stages/844/implement
-------------------------------------------------------------------------------
[Control statements](https://hyperskill.org/learn/topic/1728)
[Slices](https://hyperskill.org/learn/topic/1672)
[Working with slices](https://hyperskill.org/learn/topic/1701)
[Structs](https://hyperskill.org/learn/topic/1768)
[Parsing data from strings](https://hyperskill.org/learn/topic/1955)

##### PENDING TO WRITE TOPICS #####
-------------------------------------------------------------------------------
[Sorting](**PENDING**)
-------------------------------------------------------------------------------
###########################

##### POSSIBLE NEW TOPICS❓ #####
===============================================================================
[Manipulating Strings⁉️ 👈😆👉💯](**PENDING**) || topic about `strings` package⁉️
[Advanced Input Operations️️⁉️](**PENDING**) || topic about `bufio` and `scanner`❓
===============================================================================

*/

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Applicant struct {
	name  string
	score float64
}

func main() {
	var nApplicants int
	fmt.Scanln(&nApplicants)

	var mApplicants int
	fmt.Scanln(&mApplicants)

	var data string
	var s []string

	var name string
	var score float64

	var applicantList []Applicant

	for i := 0; i < nApplicants; i++ {
		// create a new scanner
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		// scan the line into data and separate the data into words
		data = scanner.Text()
		s = strings.Split(data, " ")

		// save the name of the applicant and the score
		name = s[0] + " " + s[1]
		score, _ = strconv.ParseFloat(s[2], 64)

		applicantList = append(applicantList, Applicant{name, score})
	}

	// sort the applicantList by highest score
	sort.Slice(applicantList, func(i, j int) bool {
		return applicantList[i].score > applicantList[j].score
	})

	fmt.Println("Successful applicants:")
	for i := 0; i < mApplicants; i++ {
		fmt.Println(applicantList[i].name)
	}
}
