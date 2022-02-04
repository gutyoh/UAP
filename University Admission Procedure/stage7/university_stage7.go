package main

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

type student []struct {
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

// function to get the max scores and max index from the scores slice of the s struct
// returns the max score and the index of the max score
func getMaxScore(s []float64) (float64, int) {
	var maxScore float64
	var maxIndex int
	for i, v := range s {
		if i == 0 {
			maxScore = v
			maxIndex = i
		} else if v > maxScore {
			maxScore = v
			maxIndex = i
		}
	}
	return maxScore, maxIndex
}

func main() {
	var nApplicants int
	fmt.Scanln(&nApplicants)

	var s student

	count := map[string]int{}
	final := map[string][]string{}

	var used []string

	file, err := os.Open("C:\\Users\\mrgut\\Documents\\UAP\\University Admission Procedure\\stage7\\applicant_list_7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		name := strings.Split(line, " ")[0]
		lastName := strings.Split(line, " ")[1]

		physScore0, _ := strconv.ParseFloat(strings.Split(line, " ")[2], 64)
		chemScore1, _ := strconv.ParseFloat(strings.Split(line, " ")[3], 64)
		mathScore2, _ := strconv.ParseFloat(strings.Split(line, " ")[4], 64)
		phyScore3, _ := strconv.ParseFloat(strings.Split(line, " ")[5], 64)
		specialScore4, _ := strconv.ParseFloat(strings.Split(line, " ")[6], 64)

		scores := []float64{physScore0, chemScore1, mathScore2, phyScore3, specialScore4}

		departments := strings.Split(line, " ")[7:]

		s = append(s, struct {
			name, lastName string
			score          []float64
			departments    []string
		}{name, lastName, scores, departments})
	}

	// --------------------------------------------------
	// First round of selection
	// Sort the students by score and first name -> for the Biotech department
	// The biotech department is index 1 within the score slice
	//sort.Slice(s, func(i, j int) bool {
	//	if (s[i].score[0]+s[i].score[1])/2 != (s[j].score[0]+s[j].score[1])/2 && (s[i].score[0]+s[i].score[1])/2 != s[j].score[4] {
	//		return (s[i].score[0]+s[i].score[1])/2 > (s[j].score[0]+s[j].score[1])/2
	//	}
	//	return s[i].name < s[j].name
	//})

	//sort.Slice(s, func(i, j int) bool {
	//	if (s[i].score[0]+s[i].score[1])/2 != (s[j].score[0]+s[j].score[1])/2 {
	//		if (s[i].score[0]+s[i].score[1])/2 != s[j].score[4] {
	//			return (s[i].score[0]+s[i].score[1])/2 > (s[j].score[0]+s[j].score[1])/2
	//		}
	//		return (s[i].score[0]+s[i].score[1])/2 > s[j].score[4]
	//	}
	//	return s[i].name < s[j].name
	//})

	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max((s[i].score[0]+s[i].score[1])/2, s[i].score[4])
		maxScoreJ := math.Max((s[j].score[0]+s[j].score[1])/2, s[j].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	// Add the students to the final map in the Biotech department
	// Iterate over all the students
	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[0] == "Biotech" && count["Biotech"] < nApplicants {
			//biotechScore := strconv.FormatFloat((s[i].score[0]+s[i].score[1])/2, 'f', 2, 64)
			//specialScore := strconv.FormatFloat(s[i].score[4], 'f', 2, 64)
			fMaxScore := math.Max((s[i].score[0]+s[i].score[1])/2, s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Biotech"] = append(final["Biotech"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Biotech"]++
		}
	}

	// Sort the students by score and first name -> for the Chemistry department
	// The biotech department is index 1 within the score slice
	//sort.Slice(s, func(i, j int) bool {
	//	if s[i].score[1] != s[j].score[1] && s[i].score[1] != s[j].score[4] {
	//		return s[i].score[1] > s[j].score[1]
	//	}
	//	return s[i].name < s[j].name
	//})
	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max(s[i].score[1], s[i].score[4])
		maxScoreJ := math.Max(s[j].score[1], s[j].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[0] == "Chemistry" && count["Chemistry"] < nApplicants {
			// chemScore := strconv.FormatFloat(s[i].score[1], 'f', 2, 64)

			fMaxScore := math.Max(s[i].score[1], s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Chemistry"] = append(final["Chemistry"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Chemistry"]++
		}
	}

	// Sort the students by score and first name -> for the Engineering department
	// The Engineering department is index 3 within the score slice
	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max((s[i].score[3]+s[i].score[2])/2, s[i].score[4])
		maxScoreJ := math.Max((s[j].score[3]+s[j].score[2])/2, s[j].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	// Add the students to the final map in the Engineering department
	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[0] == "Engineering" && count["Engineering"] < nApplicants {
			// engScore := strconv.FormatFloat((s[i].score[3]+s[i].score[2])/2, 'f', 2, 64)

			fMaxScore := math.Max((s[i].score[3]+s[i].score[2])/2, s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Engineering"] = append(final["Engineering"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Engineering"]++
		}
	}

	// Sort the students by score and first name -> for the Mathematics department
	// The Mathematics department is index 2 within the score slice
	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max(s[i].score[2], s[i].score[4])
		maxScoreJ := math.Max(s[j].score[2], s[j].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	// Add the students to the final map in the Mathematics department
	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[0] == "Mathematics" && count["Mathematics"] < nApplicants {
			// mathScore := strconv.FormatFloat(s[i].score[2], 'f', 2, 64)

			fMaxScore := math.Max(s[i].score[2], s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Mathematics"] = append(final["Mathematics"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Mathematics"]++
		}
	}

	// Sort the students by (score[0] + score[2]) and first name -> for the Physics department
	// The Physics department is index 0 within the score slice
	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max((s[i].score[0]+s[i].score[2])/2, s[i].score[4])
		maxScoreJ := math.Max((s[j].score[0]+s[j].score[2])/2, s[j].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	// Add the students to the final map in the Physics department
	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[0] == "Physics" && count["Physics"] < nApplicants {
			// phyScore := strconv.FormatFloat((s[i].score[0]+s[i].score[2])/2, 'f', 2, 64)

			fMaxScore := math.Max((s[i].score[0]+s[i].score[2])/2, s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Physics"] = append(final["Physics"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Physics"]++
		}
	}
	// --------------------------------------------------

	// --------------------------------------------------
	// Second round of appending students to the final map
	// Sort the students by score and first name -> for the Biotech department
	// The biotech department is index 1 within the score slice
	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max((s[i].score[0]+s[i].score[1])/2, s[i].score[4])
		maxScoreJ := math.Max((s[j].score[0]+s[j].score[1])/2, s[j].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	// Add the students to the final map in the Biotech department
	// Iterate over all the students
	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[1] == "Biotech" && count["Biotech"] < nApplicants {
			//biotechScore := strconv.FormatFloat((s[i].score[0]+s[i].score[1])/2, 'f', 2, 64)
			//specialScore := strconv.FormatFloat(s[i].score[4], 'f', 2, 64)
			fMaxScore := math.Max((s[i].score[0]+s[i].score[1])/2, s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Biotech"] = append(final["Biotech"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Biotech"]++
		}
	}

	// Sort the students by score and first name -> for the Chemistry department
	// The biotech department is index 1 within the score slice
	//sort.Slice(s, func(i, j int) bool {
	//	if s[i].score[1] != s[j].score[1] {
	//		if s[i].score[1] != s[j].score[4] {
	//			return s[i].score[1] > s[j].score[1]
	//		}
	//		return s[i].score[4] > s[j].score[4]
	//	}
	//	return s[i].name < s[j].name
	//})
	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max(s[i].score[1], s[i].score[4])
		maxScoreJ := math.Max(s[j].score[1], s[i].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[1] == "Chemistry" && count["Chemistry"] < nApplicants {
			// chemScore := strconv.FormatFloat(s[i].score[1], 'f', 2, 64)

			fMaxScore := math.Max(s[i].score[1], s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Chemistry"] = append(final["Chemistry"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Chemistry"]++
		}
	}

	// Sort the students by score and first name -> for the Engineering department
	// The Engineering department is index 3 within the score slice
	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max((s[i].score[3]+s[i].score[2])/2, s[i].score[4])
		maxScoreJ := math.Max((s[j].score[3]+s[j].score[2])/2, s[j].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	// Add the students to the final map in the Engineering department
	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[1] == "Engineering" && count["Engineering"] < nApplicants {
			// engScore := strconv.FormatFloat((s[i].score[3]+s[i].score[2])/2, 'f', 2, 64)

			fMaxScore := math.Max((s[i].score[3]+s[i].score[2])/2, s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Engineering"] = append(final["Engineering"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Engineering"]++
		}
	}

	// Sort the students by score and first name -> for the Mathematics department
	// The Mathematics department is index 2 within the score slice
	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max(s[i].score[2], s[i].score[4])
		maxScoreJ := math.Max(s[j].score[2], s[i].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	// Add the students to the final map in the Mathematics department
	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[1] == "Mathematics" && count["Mathematics"] < nApplicants {
			// mathScore := strconv.FormatFloat(s[i].score[2], 'f', 2, 64)

			fMaxScore := math.Max(s[i].score[2], s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Mathematics"] = append(final["Mathematics"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Mathematics"]++
		}
	}

	// Sort the students by (score[0] + score[2]) and first name -> for the Physics department
	// The Physics department is index 0 within the score slice
	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max((s[i].score[0]+s[i].score[2])/2, s[i].score[4])
		maxScoreJ := math.Max((s[j].score[0]+s[j].score[2])/2, s[j].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	// Add the students to the final map in the Physics department
	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[1] == "Physics" && count["Physics"] < nApplicants {
			// phyScore := strconv.FormatFloat((s[i].score[0]+s[i].score[2])/2, 'f', 2, 64)

			fMaxScore := math.Max((s[i].score[0]+s[i].score[2])/2, s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Physics"] = append(final["Physics"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Physics"]++
		}
	}
	// --------------------------------------------------

	// --------------------------------------------------
	// Third round of appending students to the final map
	// Sort the students by score and first name -> for the Biotech department
	// The Biotech department is index 1 within the score slice
	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max((s[i].score[0]+s[i].score[1])/2, s[i].score[4])
		maxScoreJ := math.Max((s[j].score[0]+s[j].score[1])/2, s[j].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	// Add the students to the final map in the Biotech department
	// Iterate over all the students
	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[2] == "Biotech" && count["Biotech"] < nApplicants {
			//biotechScore := strconv.FormatFloat((s[i].score[0]+s[i].score[1])/2, 'f', 2, 64)
			//specialScore := strconv.FormatFloat(s[i].score[4], 'f', 2, 64)
			fMaxScore := math.Max((s[i].score[0]+s[i].score[1])/2, s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Biotech"] = append(final["Biotech"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Biotech"]++
		}
	}

	// Sort the students by score and first name -> for the Chemistry department
	// The Chemistry department is index 1 within the score slice
	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max(s[i].score[1], s[i].score[4])
		maxScoreJ := math.Max(s[j].score[1], s[i].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[2] == "Chemistry" && count["Chemistry"] < nApplicants {
			// chemScore := strconv.FormatFloat(s[i].score[1], 'f', 2, 64)

			fMaxScore := math.Max(s[i].score[1], s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Chemistry"] = append(final["Chemistry"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Chemistry"]++
		}
	}

	// Sort the students by score and first name -> for the Engineering department
	// The Engineering department is index 3 within the score slice
	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max((s[i].score[3]+s[i].score[2])/2, s[i].score[4])
		maxScoreJ := math.Max((s[j].score[3]+s[j].score[2])/2, s[j].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	// Add the students to the final map in the Engineering department
	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[2] == "Engineering" && count["Engineering"] < nApplicants {
			// engScore := strconv.FormatFloat((s[i].score[3]+s[i].score[2])/2, 'f', 2, 64)

			fMaxScore := math.Max((s[i].score[3]+s[i].score[2])/2, s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Engineering"] = append(final["Engineering"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Engineering"]++
		}
	}

	// Sort the students by score and first name -> for the Mathematics department
	// The Mathematics department is index 2 within the score slice
	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max(s[i].score[2], s[i].score[4])
		maxScoreJ := math.Max(s[j].score[2], s[i].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	// Add the students to the final map in the Mathematics department
	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[2] == "Mathematics" && count["Mathematics"] < nApplicants {
			// mathScore := strconv.FormatFloat(s[i].score[2], 'f', 2, 64)

			fMaxScore := math.Max(s[i].score[2], s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Mathematics"] = append(final["Mathematics"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Mathematics"]++
		}
	}

	// Sort the students by (score[0] + score[2]) and first name -> for the Physics department
	// The Physics department is index 0 within the score slice
	sort.Slice(s, func(i, j int) bool {
		maxScoreI := math.Max((s[i].score[0]+s[i].score[2])/2, s[i].score[4])
		maxScoreJ := math.Max((s[j].score[0]+s[j].score[2])/2, s[j].score[4])

		if maxScoreI != maxScoreJ {
			return maxScoreI > maxScoreJ
		}
		return s[i].name < s[j].name
	})

	// Add the students to the final map in the Physics department
	for i := 0; i < len(s); i++ {
		if !isInSlice(used, s[i].name) && s[i].departments[2] == "Physics" && count["Physics"] < nApplicants {
			// phyScore := strconv.FormatFloat((s[i].score[0]+s[i].score[2])/2, 'f', 2, 64)

			fMaxScore := math.Max((s[i].score[0]+s[i].score[2])/2, s[i].score[4])

			sMaxScore := strconv.FormatFloat(fMaxScore, 'f', 2, 64)

			final["Physics"] = append(final["Physics"], s[i].name+" "+s[i].lastName+" "+sMaxScore)
			used = append(used, s[i].name)
			count["Physics"]++
		}
	}
	// --------------------------------------------------

	// ****************************************************
	// Sort each of the departments based on score and name
	sort.Slice(final["Biotech"], func(i, j int) bool {
		// sort Biotech based on score
		if strings.Split(final["Biotech"][i], " ")[2] != strings.Split(final["Biotech"][j], " ")[2] {
			return strings.Split(final["Biotech"][i], " ")[2] > strings.Split(final["Biotech"][j], " ")[2]
		}
		return strings.Split(final["Biotech"][i], " ")[0] < strings.Split(final["Biotech"][j], " ")[0]
	})

	sort.Slice(final["Chemistry"], func(i, j int) bool {
		if strings.Split(final["Chemistry"][i], " ")[2] != strings.Split(final["Chemistry"][j], " ")[2] {
			return strings.Split(final["Chemistry"][i], " ")[2] > strings.Split(final["Chemistry"][j], " ")[2]
		}
		return strings.Split(final["Chemistry"][i], " ")[0] < strings.Split(final["Chemistry"][j], " ")[0]
	})

	sort.Slice(final["Engineering"], func(i, j int) bool {
		if strings.Split(final["Engineering"][i], " ")[2] != strings.Split(final["Engineering"][j], " ")[2] {
			return strings.Split(final["Engineering"][i], " ")[2] > strings.Split(final["Engineering"][j], " ")[2]
		}
		return strings.Split(final["Engineering"][i], " ")[0] < strings.Split(final["Engineering"][j], " ")[0]
	})

	sort.Slice(final["Mathematics"], func(i, j int) bool {
		if strings.Split(final["Mathematics"][i], " ")[2] != strings.Split(final["Mathematics"][j], " ")[2] {
			return strings.Split(final["Mathematics"][i], " ")[2] > strings.Split(final["Mathematics"][j], " ")[2]
		}
		return strings.Split(final["Mathematics"][i], " ")[0] < strings.Split(final["Mathematics"][j], " ")[0]
	})

	sort.Slice(final["Physics"], func(i, j int) bool {
		if strings.Split(final["Physics"][i], " ")[2] != strings.Split(final["Physics"][j], " ")[2] {
			return strings.Split(final["Physics"][i], " ")[2] > strings.Split(final["Physics"][j], " ")[2]
		}
		return strings.Split(final["Physics"][i], " ")[0] < strings.Split(final["Physics"][j], " ")[0]
	})
	// ****************************************************

	// Print the output - Biotech first then Chem, Eng, Math, Physics
	fmt.Println("Biotech")
	file, err = os.Create("biotech.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range final["Biotech"] {
		fmt.Println(v)
		fmt.Fprintln(file, v)
	}
	fmt.Println()

	fmt.Println("Chemistry")
	file, err = os.Create("chemistry.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range final["Chemistry"] {
		fmt.Println(v)
		fmt.Fprintln(file, v)
	}
	fmt.Println()

	fmt.Println("Engineering")
	file, err = os.Create("engineering.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range final["Engineering"] {
		fmt.Println(v)
		fmt.Fprintln(file, v)
	}
	fmt.Println()

	fmt.Println("Mathematics")
	file, err = os.Create("mathematics.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range final["Mathematics"] {
		fmt.Println(v)
		fmt.Fprintln(file, v)
	}
	fmt.Println()

	fmt.Println("Physics")
	file, err = os.Create("physics.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range final["Physics"] {
		fmt.Println(v)
		fmt.Fprintln(file, v)
	}
}
