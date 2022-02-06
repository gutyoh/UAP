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
	scores   []float64
}

type ApplicantPreferences struct {
	Applicant
	departments []string
}

func contains(s []string, e string) bool {
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
			// Get the scores of the two students
			scoreI, _ := strconv.ParseFloat(strings.Split(v[i], " ")[2], 64)
			scoreJ, _ := strconv.ParseFloat(strings.Split(v[j], " ")[2], 64)

			// Get the names of the two students
			nameI := strings.Split(v[i], " ")[0]
			nameJ := strings.Split(v[j], " ")[0]

			// Sort by score first; if scores are equal, sort by first name alphabetically
			if scoreI != scoreJ {
				return scoreI > scoreJ
			}
			return nameI < nameJ
		})
	}
}

// The readApplicantPreferences func reads the applicant's preferences from the
// input file and returns a slice of ApplicantPreferences with the data of the applicants
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
			Applicant{parts[0] + " " + parts[1], scores}, parts[6:],
		})
	}
	return a
}

// The chooseFaculty function checks if the a[k].fullName is in the 'used' slice and
// if the first department of a[k].departments is the same as orderedDepartments[j] or 'dep'
// and if the count[orderedDepartments[j]] ('dep') is less than nApplicants.
func chooseFaculty(a []ApplicantPreferences, orderedDepartments []string,
	used []string, count map[string]int, final map[string][]string, nApplicants int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < len(orderedDepartments); j++ {
			dep := orderedDepartments[j]
			// Call the sortByDept function to sort students by highest score then by name alphabetically
			sortByDept(a, dep)
			switch dep {
			case "Biotech":
				for k := 0; k < len(a); k++ {
					if !contains(used, a[k].fullName) && a[k].departments[i] == dep && count[dep] < nApplicants {
						bioScore := strconv.FormatFloat(a[k].scores[1], 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+bioScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			case "Chemistry":
				for k := 0; k < len(a); k++ {
					if !contains(used, a[k].fullName) && a[k].departments[i] == dep && count[dep] < nApplicants {
						chemScore := strconv.FormatFloat(a[k].scores[1], 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+chemScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			case "Engineering":
				for k := 0; k < len(a); k++ {
					if !contains(used, a[k].fullName) && a[k].departments[i] == dep && count[dep] < nApplicants {
						engScore := strconv.FormatFloat(a[k].scores[3], 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+engScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			case "Mathematics":
				for k := 0; k < len(a); k++ {
					if !contains(used, a[k].fullName) && a[k].departments[i] == dep && count[dep] < nApplicants {
						mathScore := strconv.FormatFloat(a[k].scores[2], 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+mathScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			case "Physics":
				for k := 0; k < len(a); k++ {
					if !contains(used, a[k].fullName) && a[k].departments[i] == dep && count[dep] < nApplicants {
						phyScore := strconv.FormatFloat(a[k].scores[0], 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+phyScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			}
		}
	}
}

// ##### SORTING FUNCTION #####
/* The sortByDept function allows us to sort each of the applicant by its highest score
And then by its name alphabetically.
We start sorting with the Biotech department first then by Chem -> Eng -> Math and finish with Physics */
// -------------------------------------------------------------------------------
func sortByDept(a []ApplicantPreferences, dep string) {
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
}

// -------------------------------------------------------------------------------
// ##### END OF SORTING FUNCTION #####

func main() {
	var nApplicants int
	fmt.Scanln(&nApplicants)

	count := map[string]int{}
	final := map[string][]string{}

	var used []string

	file, err := os.Open("./applicant_list_5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// We call readApplicantPreferences to read the applicant data into 'a'
	a := readApplicantPreferences(file)

	/* The chooseFaculty function will sort the applicants starting with the Biotech dept
	 And finishing with the Physics department.
	When we sort the applicants within the Biotech dept we will first sort by the highest score
	And then by the names alphabetically. (from A to Z) */
	// chooseFaculty(a, used, count, final, nApplicants)
	chooseFaculty(a, orderedDepartments, used, count, final, nApplicants)

	// The sortApplicants function sorts the 'final' map by the highest score first and then by the names alphabetically.
	sortApplicants(final)

	// Finally, we print the applicants name and score in the alphabetical order of departments:
	// We start with the Biotech department first then Chemistry -> Engineering -> Mathematics and end with Physics.
	for i := 0; i < len(orderedDepartments); i++ {
		fmt.Println(orderedDepartments[i])
		for _, v := range final[orderedDepartments[i]] {
			fmt.Println(v)
		}
		fmt.Println()
	}
}
