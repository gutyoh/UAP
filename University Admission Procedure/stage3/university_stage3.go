package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var nApplicants int
	fmt.Scanln(&nApplicants)

	var mApplicants int
	fmt.Scanln(&mApplicants)

	var data string
	var s []string

	var name string
	var score float64

	var applicantList []interface{}

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

		// append the name and score to the applicantList as a slice of interface{}
		applicantList = append(applicantList, []interface{}{name, score})
	}

	// sort the applicantList by highest score
	sort.Slice(applicantList, func(i, j int) bool {
		return applicantList[i].([]interface{})[1].(float64) > applicantList[j].([]interface{})[1].(float64)
	})

	fmt.Println("Successful applicants:")
	for i := 0; i < mApplicants; i++ {
		fmt.Println(applicantList[i].([]interface{})[0].(string))
	}
}
