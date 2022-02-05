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

type Applicant []struct {
	name, lastName string
	score          []float64
	departments    []string
}

func isInSlice(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// The sortApplicants func iterates over the keys of the 'final' map and sorts
// the students by their highest score and then by their name alphabetically
func sortApplicants(final map[string][]string) {
	for _, v := range final {
		sort.Slice(v, func(i, j int) bool {
			if strings.Split(v[i], " ")[2] != strings.Split(v[j], " ")[2] {
				return strings.Split(v[i], " ")[2] > strings.Split(v[j], " ")[2]
			}
			return strings.Split(v[i], " ")[0] < strings.Split(v[j], " ")[0]
		})
	}
}

// The addApplicant checks if the a[i].name is in the 'used' slice and
// if the first department of a[i].departments is the same as orderedDepartments[j]
// and if the count[orderedDepartments[j]] is less than nApplicants.
func addApplicant(a Applicant, used []string, count map[string]int, final map[string][]string, nApplicants int) {
	for i := 0; i < 3; i++ {
		// Sort 'a' (applicants) by Chemistry exam score then add to "Biotech" department
		sortByChemScore(a)
		for j := 0; j < len(a); j++ {
			if !isInSlice(used, a[j].name) && a[j].departments[i] == "Biotech" && count["Biotech"] < nApplicants {
				bioScore := strconv.FormatFloat(a[j].score[1], 'f', 2, 64)

				final["Biotech"] = append(final["Biotech"], a[j].name+" "+a[j].lastName+" "+bioScore)
				used = append(used, a[j].name)
				count["Biotech"]++
			}
		}
		// Since we already sorted the applicants by Chemistry score, we can add to "Chemistry" department
		for j := 0; j < len(a); j++ {
			if !isInSlice(used, a[j].name) && a[j].departments[i] == "Chemistry" && count["Chemistry"] < nApplicants {
				chemScore := strconv.FormatFloat(a[j].score[1], 'f', 2, 64)

				final["Chemistry"] = append(final["Chemistry"], a[j].name+" "+a[j].lastName+" "+chemScore)
				used = append(used, a[j].name)
				count["Chemistry"]++
			}
		}
		// Sort the applicants by CS (Engineering) exam score then add to "Engineering" department
		sortByEngScore(a)
		for j := 0; j < len(a); j++ {
			if !isInSlice(used, a[j].name) && a[j].departments[i] == "Engineering" && count["Engineering"] < nApplicants {
				engScore := strconv.FormatFloat(a[j].score[3], 'f', 2, 64)

				final["Engineering"] = append(final["Engineering"], a[j].name+" "+a[j].lastName+" "+engScore)
				used = append(used, a[j].name)
				count["Engineering"]++
			}
		}
		// Sort the applicants by Mathematics exam score then add to "Mathematics" department
		sortByMathScore(a)
		for j := 0; j < len(a); j++ {
			if !isInSlice(used, a[j].name) && a[j].departments[i] == "Mathematics" && count["Mathematics"] < nApplicants {
				mathScore := strconv.FormatFloat(a[j].score[2], 'f', 2, 64)

				final["Mathematics"] = append(final["Mathematics"], a[j].name+" "+a[j].lastName+" "+mathScore)
				used = append(used, a[j].name)
				count["Mathematics"]++
			}
		}
		// Sort the applicants by Physics exam score then add to "Physics" department
		sortByPhyScore(a)
		for j := 0; j < len(a); j++ {
			if !isInSlice(used, a[j].name) && a[j].departments[i] == "Physics" && count["Physics"] < nApplicants {
				phyScore := strconv.FormatFloat(a[j].score[0], 'f', 2, 64)

				final["Physics"] = append(final["Physics"], a[j].name+" "+a[j].lastName+" "+phyScore)
				used = append(used, a[j].name)
				count["Physics"]++
			}
		}
	}
}

func sortByChemScore(a Applicant) {
	sort.Slice(a, func(i, j int) bool {
		if a[i].score[1] != a[j].score[1] {
			return a[i].score[1] > a[j].score[1]
		}
		return a[i].name < a[j].name
	})
}

func sortByEngScore(a Applicant) {
	sort.Slice(a, func(i, j int) bool {
		if a[i].score[3] != a[j].score[3] {
			return a[i].score[3] > a[j].score[3]
		}
		return a[i].name < a[j].name
	})
}

func sortByMathScore(a Applicant) {
	sort.Slice(a, func(i, j int) bool {
		if a[i].score[2] != a[j].score[2] {
			return a[i].score[2] > a[j].score[2]
		}
		return a[i].name < a[j].name
	})
}

func sortByPhyScore(a Applicant) {
	sort.Slice(a, func(i, j int) bool {
		if a[i].score[0] != a[j].score[0] {
			return a[i].score[0] > a[j].score[0]
		}
		return a[i].name < a[j].name
	})
}

func main() {
	var nApplicants int
	fmt.Scanln(&nApplicants)

	var a Applicant

	count := map[string]int{}
	final := map[string][]string{}

	var used []string

	file, err := os.Open("/Users/guty/PycharmProjects/UAP/University Admission Procedure/stage5/applicant_list_5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		name := strings.Split(line, " ")[0]
		lastName := strings.Split(line, " ")[1]

		phyScore, _ := strconv.ParseFloat(strings.Split(line, " ")[2], 64)
		chemScore, _ := strconv.ParseFloat(strings.Split(line, " ")[3], 64)
		mathScore, _ := strconv.ParseFloat(strings.Split(line, " ")[4], 64)
		engScore, _ := strconv.ParseFloat(strings.Split(line, " ")[5], 64)

		scores := []float64{phyScore, chemScore, mathScore, engScore}

		departments := strings.Split(line, " ")[6:]

		a = append(a, struct {
			name, lastName string
			score          []float64
			departments    []string
		}{name, lastName, scores, departments})
	}

	/* The addApplicant function will sort the applicants starting with the Biotech dept
	 And finishing with the Physics department.
	When we sort the applicants within the Biotech dept we will first sort by the highest score
	And then by the names alphabetically. (from A to Z) */
	addApplicant(a, used, count, final, nApplicants)

	// The sortApplicants function sorts the 'final' map by the highest score first and then by the names alphabetically.
	sortApplicants(final)

	// Finally, we print the output - The Biotech department applicants first then Chem, Eng, Math, Physics
	fmt.Println("Biotech")
	for _, v := range final["Biotech"] {
		fmt.Println(v)
	}
	fmt.Println()

	fmt.Println("Chemistry")
	for _, v := range final["Chemistry"] {
		fmt.Println(v)
	}
	fmt.Println()

	fmt.Println("Engineering")
	for _, v := range final["Engineering"] {
		fmt.Println(v)
	}
	fmt.Println()

	fmt.Println("Mathematics")
	for _, v := range final["Mathematics"] {
		fmt.Println(v)
	}
	fmt.Println()

	fmt.Println("Physics")
	for _, v := range final["Physics"] {
		fmt.Println(v)
	}
}
