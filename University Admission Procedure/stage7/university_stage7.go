package main

/*
[University Admission Procedure - Stage 7/7: Something special](https://hyperskill.org/projects/163/stages/850/implement)
-------------------------------------------------------------------------------
##### 🚫 NO NEW TOPICS REQUIRED 🚫 #####
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

type FinalApplicant struct {
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

		// Here we create a new variable 'specialScore' to add the new **special score**!
		specialScore, _ := strconv.ParseFloat(parts[6], 64)

		scores := []float64{phyScore, chemScore, mathScore, engScore, specialScore}

		a = append(a, ApplicantPreferences{
			parts[0] + " " + parts[1], scores, parts[7:],
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

func sortByDept(a []ApplicantPreferences, dep string) []ApplicantPreferences {
	switch dep {
	case "Biotech":
		sort.Slice(a, func(i, j int) bool {
			maxScoreI := math.Max((a[i].scores[0]+a[i].scores[1])/2, a[i].scores[4])
			maxScoreJ := math.Max((a[j].scores[0]+a[j].scores[1])/2, a[j].scores[4])

			if maxScoreI != maxScoreJ {
				return maxScoreI > maxScoreJ
			}
			return a[i].fullName < a[j].fullName
		})
	case "Chemistry":
		sort.Slice(a, func(i, j int) bool {
			maxScoreI := math.Max(a[i].scores[1], a[i].scores[4])
			maxScoreJ := math.Max(a[j].scores[1], a[j].scores[4])

			if maxScoreI != maxScoreJ {
				return maxScoreI > maxScoreJ
			}
			return a[i].fullName < a[j].fullName
		})
	case "Engineering":
		sort.Slice(a, func(i, j int) bool {
			maxScoreI := math.Max((a[i].scores[3]+a[i].scores[2])/2, a[i].scores[4])
			maxScoreJ := math.Max((a[j].scores[3]+a[j].scores[2])/2, a[j].scores[4])

			if maxScoreI != maxScoreJ {
				return maxScoreI > maxScoreJ
			}
			return a[i].fullName < a[j].fullName
		})
	case "Mathematics":
		sort.Slice(a, func(i, j int) bool {
			maxScoreI := math.Max(a[i].scores[2], a[i].scores[4])
			maxScoreJ := math.Max(a[j].scores[2], a[j].scores[4])

			if maxScoreI != maxScoreJ {
				return maxScoreI > maxScoreJ
			}
			return a[i].fullName < a[j].fullName
		})
	case "Physics":
		sort.Slice(a, func(i, j int) bool {
			maxScoreI := math.Max((a[i].scores[0]+a[i].scores[2])/2, a[i].scores[4])
			maxScoreJ := math.Max((a[j].scores[0]+a[j].scores[2])/2, a[j].scores[4])

			if maxScoreI != maxScoreJ {
				return maxScoreI > maxScoreJ
			}
			return a[i].fullName < a[j].fullName
		})
	}
	return a
}

func main() {
	var nApplicants int
	fmt.Scanln(&nApplicants)

	file, err := os.Open("./applicant_list_7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	applicants := readApplicantPreferences(file)

	exam := map[string][]int{
		"Biotech":     {0, 1, 4},
		"Chemistry":   {1, 1, 4},
		"Engineering": {2, 3, 4},
		"Mathematics": {2, 2, 4},
		"Physics":     {0, 2, 4},
	}

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
			applicantsSorted := sortByDept(applicants, dep)
			for _, applicant := range applicantsSorted {
				if applicant.departments[i] == dep && len(departments[dep]) < nApplicants && !contains(used, applicant.fullName) {
					score := math.Max((applicant.scores[exam[dep][0]]+applicant.scores[exam[dep][1]])/2, applicant.scores[exam[dep][2]])

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
		file, err = os.Create(strings.ToLower(dep) + ".txt")
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range departments[dep] {
			fmt.Printf("%s %.2f\n", v.fullName, v.score)
			_, err = fmt.Fprintf(file, "%s %.2f\n", v.fullName, v.score)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println()
	}
}
