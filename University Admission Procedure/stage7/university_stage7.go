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
			scoreI := strings.Split(v[i], " ")[2]
			scoreJ := strings.Split(v[j], " ")[2]

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

		// Here we create a new variable 'specialScore' to add the new **special score**!
		specialScore, _ := strconv.ParseFloat(parts[6], 64)

		scores := []float64{phyScore, chemScore, mathScore, engScore, specialScore}

		a = append(a, ApplicantPreferences{
			Applicant{parts[0] + " " + parts[1], scores}, parts[7:],
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
						specialScore := a[k].scores[4]

						// Get the highest score between the average score of the
						// [(Chemistry + Physics)/2] exams and the 'Special' exam
						bioMaxScore := math.Max(avgScore, specialScore)

						finalMaxScore := strconv.FormatFloat(bioMaxScore, 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+finalMaxScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			case "Chemistry":
				for k := 0; k < len(a); k++ {
					if !contains(used, a[k].fullName) && a[k].departments[i] == dep && count[dep] < nApplicants {
						// Get the highest score between Chemistry exam and the 'Special' exam
						chemMaxScore := math.Max(a[k].scores[1], a[k].scores[4])

						finalMaxScore := strconv.FormatFloat(chemMaxScore, 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+finalMaxScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			case "Engineering":
				for k := 0; k < len(a); k++ {
					if !contains(used, a[k].fullName) && a[k].departments[i] == dep && count[dep] < nApplicants {
						avgScore := (a[k].scores[3] + a[k].scores[2]) / 2
						specialScore := a[k].scores[4]

						// Get the highest score between the average score of the
						// [(Engineering + Mathematics)/2] exams and the 'Special' exam
						engMaxScore := math.Max(avgScore, specialScore)

						finalMaxScore := strconv.FormatFloat(engMaxScore, 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+finalMaxScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			case "Mathematics":
				for k := 0; k < len(a); k++ {
					if !contains(used, a[k].fullName) && a[k].departments[i] == dep && count[dep] < nApplicants {
						// Get the highest score between Mathematics exam and the 'Special' exam
						mathMaxScore := math.Max(a[k].scores[2], a[k].scores[4])

						finalMaxScore := strconv.FormatFloat(mathMaxScore, 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+finalMaxScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			case "Physics":
				for k := 0; k < len(a); k++ {
					if !contains(used, a[k].fullName) && a[k].departments[i] == dep && count[dep] < nApplicants {
						avgScore := (a[k].scores[0] + a[k].scores[2]) / 2
						specialScore := a[k].scores[4]

						// Get the highest score between the average score of the
						// [(Engineering + Mathematics)/2] exams and the 'Special' exam
						phyMaxScore := math.Max(avgScore, specialScore)

						finalMaxScore := strconv.FormatFloat(phyMaxScore, 'f', 2, 64)

						final[dep] = append(final[dep], a[k].fullName+" "+finalMaxScore)
						used = append(used, a[k].fullName)
						count[dep]++
					}
				}
			}
		}
	}
}

/* The sorting functions are updated to properly calculate the highest score
The single score of one dept. or average of score of two depts. is compared to the special score
Then we get the max score between the avg. score and the special score
And finally we sort the applicants by the highest max score and then by the name alphabetically */

// ##### SORTING FUNCTIONS #####
// -------------------------------------------------------------------------------

func sortByDept(a []ApplicantPreferences, dep string) {
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

	file, err := os.Open("./applicant_list_7.txt")
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
