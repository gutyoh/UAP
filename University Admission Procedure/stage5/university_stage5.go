package main

/*
[University Admission Procedure - Stage 5/7: Special knowledge](https://hyperskill.org/projects/163/stages/848/implement)
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

type Applicant struct {
	fullName string
	score    float64
}

type ApplicantPreferences struct {
	fullName    string
	scores      []float64
	departments []string
}

func readApplicantPreferences(file *os.File) []ApplicantPreferences {
	var a []ApplicantPreferences
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		phyScore, _ := strconv.ParseFloat(parts[2], 64)
		chemScore, _ := strconv.ParseFloat(parts[3], 64)
		mathScore, _ := strconv.ParseFloat(parts[4], 64)
		engScore, _ := strconv.ParseFloat(parts[5], 64)

		scores := []float64{phyScore, chemScore, mathScore, engScore}

		a = append(a, ApplicantPreferences{
			parts[0] + " " + parts[1], scores, parts[6:],
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

func chooseFaculty(applicants []ApplicantPreferences, nApplicants int, departments map[string][]Applicant, exam map[string][]int, used []string) {
	for i := 0; i < 3; i++ {
		for _, dep := range orderedDepartments {
			applicantsSorted := sortByDept(applicants, dep)
			for _, applicant := range applicantsSorted {
				if applicant.departments[i] == dep && len(departments[dep]) < nApplicants && !contains(used, applicant.fullName) {
					score := applicant.scores[exam[dep][0]]

					departments[dep] = append(departments[dep], Applicant{applicant.fullName, score})

					used = append(used, applicant.fullName)
				}
			}
		}
	}
}

func sortByDept(a []ApplicantPreferences, dep string) []ApplicantPreferences {
	switch dep {
	case "Biotech":
		sort.Slice(a, func(i, j int) bool {
			if a[i].scores[1] != a[j].scores[1] {
				return a[i].scores[1] > a[j].scores[1]
			}
			return a[i].fullName < a[j].fullName
		})
	case "Chemistry":
		sort.Slice(a, func(i, j int) bool {
			if a[i].scores[1] != a[j].scores[1] {
				return a[i].scores[1] > a[j].scores[1]
			}
			return a[i].fullName < a[j].fullName
		})
	case "Engineering":
		sort.Slice(a, func(i, j int) bool {
			if a[i].scores[3] != a[j].scores[3] {
				return a[i].scores[3] > a[j].scores[3]
			}
			return a[i].fullName < a[j].fullName
		})
	case "Mathematics":
		sort.Slice(a, func(i, j int) bool {
			if a[i].scores[2] != a[j].scores[2] {
				return a[i].scores[2] > a[j].scores[2]
			}
			return a[i].fullName < a[j].fullName
		})
	case "Physics":
		sort.Slice(a, func(i, j int) bool {
			if a[i].scores[0] != a[j].scores[0] {
				return a[i].scores[0] > a[j].scores[0]
			}
			return a[i].fullName < a[j].fullName
		})
	}
	return a
}

func prepareFinalOrder(departments map[string][]Applicant) {
	for _, dep := range orderedDepartments {
		sort.Slice(departments[dep], func(i, j int) bool {
			if departments[dep][i].score != departments[dep][j].score {
				return departments[dep][i].score > departments[dep][j].score
			}
			return departments[dep][i].fullName < departments[dep][j].fullName
		})
	}
}

func showAccepted(departments map[string][]Applicant) {
	for _, dep := range orderedDepartments {
		fmt.Println(dep)
		for _, v := range departments[dep] {
			fmt.Printf("%s %.2f\n", v.fullName, v.score)
		}
		fmt.Println()
	}
}

func main() {
	var nApplicants int
	fmt.Scanln(&nApplicants)

	file, err := os.Open("./applicant_list_5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	applicants := readApplicantPreferences(file)

	exam := map[string][]int{
		"Biotech":     {1},
		"Chemistry":   {1},
		"Engineering": {3},
		"Mathematics": {2},
		"Physics":     {0},
	}

	departments := map[string][]Applicant{
		"Biotech":     {},
		"Chemistry":   {},
		"Engineering": {},
		"Mathematics": {},
		"Physics":     {},
	}

	var used []string

	chooseFaculty(applicants, nApplicants, departments, exam, used)
	prepareFinalOrder(departments)
	showAccepted(departments)
}
