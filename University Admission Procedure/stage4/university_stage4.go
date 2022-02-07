package main

/*
[University Admission Procedure - Stage 4/7: Choose your path](https://hyperskill.org/projects/163/stages/847/implement)
-------------------------------------------------------------------------------
##### 🚫 NO NEW TOPICS REQUIRED 🚫 #####
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

type FinalApplicant struct {
	fullName string
	score    float64
}

type ApplicantPreferences struct {
	fullName    string
	score       float64
	departments []string
}

func readApplicantPreferences(file *os.File) []ApplicantPreferences {
	var a []ApplicantPreferences
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		score, _ := strconv.ParseFloat(parts[2], 64)

		a = append(a, ApplicantPreferences{
			parts[0] + " " + parts[1], score, parts[3:],
		})
	}
	return a
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func sortByDept(a []ApplicantPreferences) []ApplicantPreferences {
	sort.Slice(a, func(i, j int) bool {
		if a[i].score != a[j].score {
			return a[i].score > a[j].score
		}
		return a[i].fullName < a[j].fullName
	})
	return a
}

func main() {
	var nApplicants int
	fmt.Scanln(&nApplicants)

	file, err := os.Open("./applicant_list_4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	applicants := readApplicantPreferences(file)

	departments := map[string][]FinalApplicant{
		"Biotech":     {},
		"Chemistry":   {},
		"Engineering": {},
		"Mathematics": {},
		"Physics":     {},
	}

	var used []string

	for i := 0; i < 3; i++ {
		for _, dep := range orderedDepartments {
			applicantsSorted := sortByDept(applicants)
			for _, applicant := range applicantsSorted {
				if applicant.departments[i] == dep && len(departments[dep]) < nApplicants && !contains(used, applicant.fullName) {
					score := applicant.score

					departments[dep] = append(departments[dep], FinalApplicant{applicant.fullName, score})

					used = append(used, applicant.fullName)
				}
			}
		}
	}

	for _, dep := range orderedDepartments {
		sort.Slice(departments[dep], func(i, j int) bool {
			if departments[dep][i].score != departments[dep][j].score {
				return departments[dep][i].score > departments[dep][j].score
			}
			return departments[dep][i].fullName < departments[dep][j].fullName
		})
	}

	for _, dep := range orderedDepartments {
		fmt.Println(dep)
		for _, v := range departments[dep] {
			fmt.Printf("%s %.2f\n", v.fullName, v.score)
		}
		fmt.Println()
	}
}
