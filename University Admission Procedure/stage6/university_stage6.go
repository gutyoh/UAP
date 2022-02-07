package main

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

type Departments struct {
	depName        string
	FinalApplicant []FinalApplicant
}

type FinalApplicant struct {
	fullName string
	score    float64
}

type ApplicantPreferences struct {
	fullName    string
	scores      []float64
	departments []string
}

type Exam struct {
	depName string
	examNum []int
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

		// Here we create a new variable 'specialScore' to add the new **special score**!
		// specialScore, _ := strconv.ParseFloat(parts[6], 64)

		scores := []float64{phyScore, chemScore, mathScore, engScore}

		a = append(a, ApplicantPreferences{
			parts[0] + " " + parts[1], scores, parts[6:],
		})
	}
	return a
}

func sortByDept(a []ApplicantPreferences, dep string) []ApplicantPreferences {
	switch dep {
	case "Biotech":
		sort.Slice(a, func(i, j int) bool {
			maxScoreI := (a[i].scores[0] + a[i].scores[1]) / 2
			maxScoreJ := (a[j].scores[0] + a[j].scores[1]) / 2

			if maxScoreI != maxScoreJ {
				return maxScoreI > maxScoreJ
			}
			return strings.Split(a[i].fullName, " ")[0] < strings.Split(a[j].fullName, " ")[0]
		})
	case "Chemistry":
		sort.Slice(a, func(i, j int) bool {
			if a[i].scores[1] != a[j].scores[1] {
				return a[i].scores[1] > a[j].scores[1]
			}
			return strings.Split(a[i].fullName, " ")[0] < strings.Split(a[j].fullName, " ")[0]
		})
	case "Engineering":
		sort.Slice(a, func(i, j int) bool {
			maxScoreI := (a[i].scores[3] + a[i].scores[2]) / 2
			maxScoreJ := (a[j].scores[3] + a[j].scores[2]) / 2

			if maxScoreI != maxScoreJ {
				return maxScoreI > maxScoreJ
			}
			return strings.Split(a[i].fullName, " ")[0] < strings.Split(a[j].fullName, " ")[0]
		})
	case "Mathematics":
		sort.Slice(a, func(i, j int) bool {
			maxScoreI := a[i].scores[2]
			maxScoreJ := a[j].scores[2]

			if maxScoreI != maxScoreJ {
				return maxScoreI > maxScoreJ
			}
			return strings.Split(a[i].fullName, " ")[0] < strings.Split(a[j].fullName, " ")[0]
		})
	case "Physics":
		sort.Slice(a, func(i, j int) bool {
			maxScoreI := (a[i].scores[0] + a[i].scores[2]) / 2
			maxScoreJ := (a[j].scores[0] + a[j].scores[2]) / 2

			if maxScoreI != maxScoreJ {
				return maxScoreI > maxScoreJ
			}
			return strings.Split(a[i].fullName, " ")[0] < strings.Split(a[j].fullName, " ")[0]
		})
	}
	return a
}

func removeApplicant(a []ApplicantPreferences, fullName string) []ApplicantPreferences {
	for i, v := range a {
		if v.fullName == fullName {
			a = append(a[:i], a[i+1:]...)
			return a
		}
	}
	return a
}

func main() {
	var nApplicants int
	fmt.Scanln(&nApplicants)

	file, err := os.Open("/Users/guty/PycharmProjects/UAP/University Admission Procedure/stage6/applicant_list_6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	applicants := readApplicantPreferences(file)

	exam := map[string][]int{
		"Biotech":     {0, 1},
		"Chemistry":   {1, 1},
		"Engineering": {2, 3},
		"Mathematics": {2, 2},
		"Physics":     {0, 2},
	}

	departments := map[string][]FinalApplicant{
		"Biotech":     {},
		"Chemistry":   {},
		"Engineering": {},
		"Mathematics": {},
		"Physics":     {},
	}

	for i := 0; i < 3; i++ {
		for _, dep := range orderedDepartments {
			applicantsSorted := sortByDept(applicants, dep)
			for _, applicant := range applicantsSorted {
				if applicant.departments[i] == dep && len(departments[dep]) < nApplicants {
					score := (applicant.scores[exam[dep][0]] + applicant.scores[exam[dep][1]]) / 2

					departments[dep] = append(departments[dep], FinalApplicant{applicant.fullName, score})

					// remove 'applicant' from 'applicants' - esto esta borrando a applicants y por eso el sort falla
					applicants = removeApplicant(applicants, applicant.fullName)
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
		file, err = os.Create(strings.ToLower(dep) + ".txt")
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range departments[dep] {
			fmt.Printf("%s %.2f\n", v.fullName, v.score)
			fmt.Fprintf(file, "%s %.2f\n", v.fullName, v.score)
		}
		fmt.Println()
	}
}
