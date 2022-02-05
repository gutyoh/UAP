package main

/*
University Admission Procedure - Stage 4/7: [Choose your path](https://hyperskill.org/projects/163/stages/847/implement)
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
			scoreI := strings.Split(v[i], " ")[2]
			scoreJ := strings.Split(v[j], " ")[2]

			nameI := strings.Split(v[i], " ")[0]
			nameJ := strings.Split(v[j], " ")[0]

			//if strings.Split(v[i], " ")[2] != strings.Split(v[j], " ")[2] {
			//	return strings.Split(v[i], " ")[2] > strings.Split(v[j], " ")[2]
			//}
			//return strings.Split(v[i], " ")[0] < strings.Split(v[j], " ")[0]

			if scoreI != scoreJ {
				return scoreI > scoreJ
			}
			return nameI < nameJ
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

	orderedDepartments := []string{
		"Biotech",
		"Chemistry",
		"Engineering",
		"Mathematics",
		"Physics",
	}

	// Open the applicant_list_4.txt file to read the applicants, their scores and department choices
	file, err := os.Open("C:\\Users\\mrgut\\Documents\\UAP\\University Admission Procedure\\stage4\\applicant_list_4.txt")
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
			// Create a string 'name' that contains the name, last name and score of the applicant
			name := entry.name + " " + entry.lastName + " " + strconv.FormatFloat(entry.score, 'f', 2, 64)

			// In case the applicant is within the 'used' slice, skip it otherwise append him/her to the 'final' map:
			if !isInSlice(used, name) && count[entry.departments[i]] < nApplicants {
				final[entry.departments[i]] = append(final[entry.departments[i]], name)
				used = append(used, name)
				count[entry.departments[i]]++
			}

			//if count[entry.departments[i]] == nApplicants || isInSlice(used, name) {
			//	continue
			//} else {
			//	final[entry.departments[i]] = append(final[entry.departments[i]], name)
			//	count[entry.departments[i]] += 1
			//	used = append(used, name)
			//}
		}
	}

	// Call the sortApplicants function to sort the final map by the highest score and then by the name alphabetically
	sortApplicants(final)

	// Finally, we print the applicants name and score in the alphabetical order of departments:
	// We start with the Biotech department first then Chemistry -> Engineering -> Mathematics and end with Physics.
	for i := 0; i < len(orderedDepartments); i++ {
		fmt.Println(orderedDepartments[i])
		fileName := strings.ToLower(orderedDepartments[i]) + ".txt"
		file, err = os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range final[orderedDepartments[i]] {
			fmt.Println(v)
		}
	}
}
