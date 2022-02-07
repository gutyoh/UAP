package main

/*
[University Admission Procedure - Stage 7/7: Something special](https://hyperskill.org/projects/163/stages/850/implement)
-------------------------------------------------------------------------------
[Math package] **TODO**
*/

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

type (
	Applicant struct {
		fullName string
	}

	University struct {
		applicants           []Applicant
		applicantScores      map[string][]float64
		applicantPreferences map[string][]string

		finals map[string][]ExamResult
	}

	ExamResult struct {
		Applicant
		score float64
	}
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (u *University) getApplications(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		phyScore, _ := strconv.ParseFloat(parts[2], 64)
		chemScore, _ := strconv.ParseFloat(parts[3], 64)
		mathScore, _ := strconv.ParseFloat(parts[4], 64)
		engScore, _ := strconv.ParseFloat(parts[5], 64)

		// Here we create a new variable 'specialScore' to add the new **special score**!
		specialScore, _ := strconv.ParseFloat(parts[6], 64)

		scores := []float64{phyScore, chemScore, mathScore, engScore, specialScore}
		fullName := parts[0] + " " + parts[1]

		u.applicants = append(u.applicants, Applicant{fullName})
		u.applicantScores[fullName] = scores
		u.applicantPreferences[fullName] = parts[7:]
	}
}

func (u *University) chooseFaculty(nApplicants int) {
	accepted := make([]string, 0, len(u.applicants))
	for i := 0; i < 3; i++ {
		for _, dep := range orderedDepartments {
			u.sortByMajorScore(dep)
			for _, a := range u.applicants {
				if contains(accepted, a.fullName) ||
					len(u.finals[dep]) == nApplicants ||
					u.applicantPreferences[a.fullName][i] != dep {
					continue
				}
				u.finals[dep] = append(
					u.finals[dep], ExamResult{a, u.majorScoreForDepartment(a, dep)},
				)
				accepted = append(accepted, a.fullName)
			}
		}
	}
}

func (u *University) majorScoreForDepartment(a Applicant, dep string) float64 {
	switch dep {
	case "Physics":
		return math.Max((u.applicantScores[a.fullName][0]+u.applicantScores[a.fullName][2])/2,
			u.applicantScores[a.fullName][4])
	case "Biotech":
		return math.Max((u.applicantScores[a.fullName][0]+u.applicantScores[a.fullName][1])/2,
			u.applicantScores[a.fullName][4])
	case "Mathematics":
		return math.Max(u.applicantScores[a.fullName][2], u.applicantScores[a.fullName][4])
	case "Engineering":
		return math.Max((u.applicantScores[a.fullName][3]+u.applicantScores[a.fullName][2])/2,
			u.applicantScores[a.fullName][4])
	default: // Chemistry
		return math.Max(u.applicantScores[a.fullName][1], u.applicantScores[a.fullName][4])
	}
}

func (u *University) sortByMajorScore(dep string) {
	sort.Slice(u.applicants, func(i, j int) bool {
		first, second := u.applicants[i], u.applicants[j]
		if u.majorScoreForDepartment(first, dep) != u.majorScoreForDepartment(second, dep) {
			return u.majorScoreForDepartment(first, dep) > u.majorScoreForDepartment(second, dep)
		}
		return first.fullName < second.fullName
	})
}

func (u *University) prepareFinalOrder() {
	for _, dep := range orderedDepartments {
		sort.Slice(u.finals[dep], func(i, j int) bool {
			first, second := u.finals[dep][i], u.finals[dep][j]
			if first.score != second.score {
				return first.score > second.score
			}
			return first.fullName < second.fullName
		})
	}
}

func (u *University) showAccepted() {
	for _, dep := range orderedDepartments {
		fmt.Println(dep)
		fileName := strings.ToLower(dep) + ".txt"
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range u.finals[dep] {
			fmt.Printf("%s %.2f\n", v.fullName, v.score)
			_, err = fmt.Fprintf(file, "%s %.2f\n", v.fullName, v.score)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println()
	}
}

func main() {
	var nApplicants int
	fmt.Scanln(&nApplicants)

	file, err := os.Open("./applicant_list_7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	u := University{
		applicantScores:      make(map[string][]float64),
		applicantPreferences: make(map[string][]string),
		finals:               make(map[string][]ExamResult),
	}
	u.getApplications(file)
	u.chooseFaculty(nApplicants)
	u.prepareFinalOrder()
	u.showAccepted()
}
