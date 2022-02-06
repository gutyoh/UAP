package main

/*
[University Admission Procedure - Stage 6/7: Extensive training](https://hyperskill.org/projects/163/stages/849/implement)
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

// The contains function checks if a certain string is within a slice of strings.
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
						avgScore := (a[k].scores[0] + a[k].scores[1]) / 2
						bioScore := strconv.FormatFloat(avgScore, 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+bioScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			case "Chemistry":
				for k := 0; k < len(a); k++ {
					if !contains(used, a[k].fullName) && a[k].departments[i] == dep && count[dep] < nApplicants {
						avgScore := a[k].scores[1]
						chemScore := strconv.FormatFloat(avgScore, 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+chemScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			case "Engineering":
				for k := 0; k < len(a); k++ {
					if !contains(used, a[k].fullName) && a[k].departments[i] == dep && count[dep] < nApplicants {
						avgScore := (a[k].scores[3] + a[k].scores[2]) / 2
						engScore := strconv.FormatFloat(avgScore, 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+engScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			case "Mathematics":
				for k := 0; k < len(a); k++ {
					if !contains(used, a[k].fullName) && a[k].departments[i] == dep && count[dep] < nApplicants {
						avgScore := a[k].scores[2]
						mathScore := strconv.FormatFloat(avgScore, 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+mathScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			case "Physics":
				for k := 0; k < len(a); k++ {
					if !contains(used, a[k].fullName) && a[k].departments[i] == dep && count[dep] < nApplicants {
						avgScore := (a[k].scores[0] + a[k].scores[2]) / 2
						phyScore := strconv.FormatFloat(avgScore, 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+phyScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			}
		}
	}
}

// ##### SORTING FUNCTIONS #####
// -------------------------------------------------------------------------------
/* We update and rename the sortByDept function in Stage #6 - because now we require
to calculate the average score of 2 exams for the Biotech, Engineering and Physics departments.
The Chemistry and Mathematics departments are not required to calculate the average score!
*/
func sortByDept(a []ApplicantPreferences, dep string) {
	switch dep {
	case "Biotech":
		sort.Slice(a, func(i, j int) bool {
			if (a[i].scores[0]+a[i].scores[1])/2 != (a[j].scores[0]+a[j].scores[1])/2 {
				return (a[i].scores[0]+a[i].scores[1])/2 > (a[j].scores[0]+a[j].scores[1])/2
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
			if (a[i].scores[3]+a[i].scores[2])/2 != (a[j].scores[3]+a[j].scores[2])/2 {
				return (a[i].scores[3]+a[i].scores[2])/2 > (a[j].scores[3]+a[j].scores[2])/2
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
			if (a[i].scores[0]+a[i].scores[2])/2 != (a[j].scores[0]+a[j].scores[2])/2 {
				return (a[i].scores[0]+a[i].scores[2])/2 > (a[j].scores[0]+a[j].scores[2])/2
			}
			return a[i].fullName < a[j].fullName
		})
	}
}

// -------------------------------------------------------------------------------
// ##### END OF SORTING FUNCTION #####

// The writeData func prints and writes the output to each file in order of departments.
// It starts with the Biotech department first then Chemistry -> Engineering -> Mathematics and end with Physics.
func writeData(file *os.File, orderedDepartments []string, final map[string][]string) {
	for i := 0; i < len(orderedDepartments); i++ {
		fmt.Println(orderedDepartments[i])
		fileName := strings.ToLower(orderedDepartments[i]) + ".txt"
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range final[orderedDepartments[i]] {
			fmt.Println(v)
			_, err = fmt.Fprintln(file, v)
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

	count := map[string]int{}
	final := map[string][]string{}

	var used []string

	file, err := os.Open("./applicant_list_6.txt")
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
	chooseFaculty(a, orderedDepartments, used, count, final, nApplicants)

	// The sortApplicants function sorts the 'final' map by the highest score first and then by the names alphabetically.
	sortApplicants(final)

	// Finally, we call writeData to print and write the output to each file in order of departments.
	writeData(file, orderedDepartments, final)
}
