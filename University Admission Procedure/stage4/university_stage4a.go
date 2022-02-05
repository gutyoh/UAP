package main

/*
[University Admission Procedure - Stage 4/7: Choose your path](https://hyperskill.org/projects/163/stages/847/implement)
-------------------------------------------------------------------------------
[Function decomposition](https://hyperskill.org/learn/topic/1893)
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var orderedDepartments = []string{
	"Biotech",
	"Chemistry",
	"Engineering",
	"Mathematics",
	"Physics",
}

type Applicant struct {
	fullName string
	score    float64
}

type ApplicantPreferences struct {
	Applicant
	departments []string
}

func contains(s []Applicant, e Applicant) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func readApplicantPreferences(file *os.File) []ApplicantPreferences {
	var a []ApplicantPreferences
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		score, _ := strconv.ParseFloat(parts[2], 64)

		a = append(a, ApplicantPreferences{
			Applicant{parts[0] + " " + parts[1], score},
			parts[3:],
		})
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i].score != a[j].score {
			return a[i].score > a[j].score
		}
		return a[i].fullName < a[j].fullName
	})

	return a
}

func chooseFaculty(a []ApplicantPreferences, nApplicants int) map[string][]Applicant {
	count := map[string]int{}
	final := map[string][]Applicant{}

	var used []Applicant
	for i := 0; i < 3; i++ {
		for _, entry := range a {
			if !contains(used, entry.Applicant) && count[entry.departments[i]] < nApplicants {
				final[entry.departments[i]] = append(final[entry.departments[i]], entry.Applicant)
				used = append(used, entry.Applicant)
				count[entry.departments[i]]++
			}
		}
	}

	for _, v := range final {
		sort.Slice(v, func(i, j int) bool {
			if v[i].score != v[j].score {
				return v[i].score > v[j].score
			}
			return v[i].fullName < v[j].fullName
		})
	}
	return final
}

func main() {
	var nApplicants int
	fmt.Scanln(&nApplicants)

	file, err := os.Open("./applicant_list_4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	a := readApplicantPreferences(file)
	final := chooseFaculty(a, nApplicants)

	for i := 0; i < len(orderedDepartments); i++ {
		fmt.Println(orderedDepartments[i])
		for _, v := range final[orderedDepartments[i]] {
			fmt.Printf("%s %.2f\n", v.fullName, v.score)
		}
		fmt.Println()
	}
}