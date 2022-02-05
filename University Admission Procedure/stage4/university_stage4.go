package main

/*
University Admission Procedure - Stage 4/7: Choose your path
https://hyperskill.org/projects/163/stages/844/implement
-------------------------------------------------------------------------------
[Functions](https://hyperskill.org/learn/topic/1750)
[Maps](https://hyperskill.org/learn/topic/1824)
[Working with files in Go](https://hyperskill.org/learn/topic/1768)
[Reading files in Go](https://hyperskill.org/learn/topic/1787)

##### PENDING TOPICS #####
-------------------------------------------------------------------------------
[Sorting](**PENDING**)
-------------------------------------------------------------------------------

##### POSSIBLE NEW TOPICS FOR THE GRAPH ⁉️ #####
===============================================================================
[Manipulating Strings⁉️ 👈😆👉💯](**PENDING**) || topic about `strings` package⁉️
[Advanced Input Operations️️⁉️](**PENDING**) || topic about `bufio` and `scanner`❓
===============================================================================
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

type Applicant []struct {
	name, lastName string
	score          float64
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

func main() {
	var nApplicants int
	fmt.Scanln(&nApplicants)

	var a Applicant

	count := map[string]int{}
	final := map[string][]string{}
	var used []string

	// Open the applicant_list.txt file to read the applicants, their scores and department choices
	file, err := os.Open("applicant_list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		name := strings.Split(line, " ")[0]
		lastName := strings.Split(line, " ")[1]
		score, _ := strconv.ParseFloat(strings.Split(line, " ")[2], 64)
		departments := strings.Split(line, " ")[3:]

		a = append(a, struct {
			name, lastName string
			score          float64
			departments    []string
		}{name, lastName, score, departments})
	}

	// Sort the applicants by score and first name, to add them in the correct order to the final map
	sort.Slice(a, func(i, j int) bool {
		if a[i].score != a[j].score {
			return a[i].score > a[j].score
		}
		return a[i].name < a[j].name
	})

	// Iterate over the applicants and add them to the final map
	for i := 0; i < 3; i++ {
		for _, entry := range a {
			name := entry.name + " " + entry.lastName + " " + strconv.FormatFloat(entry.score, 'f', 2, 64)
			if count[entry.departments[i]] == nApplicants || isInSlice(used, name) {
				continue
			} else {
				final[entry.departments[i]] = append(final[entry.departments[i]], name)
				count[entry.departments[i]] += 1
				used = append(used, name)
			}
		}
	}

	// Call the sortApplicants function to sort the final map by the highest score and then by the name alphabetically
	sortApplicants(final)

	// Print the output of each department - Biotech first then Chem, Eng, Math, Physics
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
