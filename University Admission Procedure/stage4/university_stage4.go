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

type student []struct {
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

func main() {
	var nApplicants int
	fmt.Scanln(&nApplicants)

	var s student

	count := map[string]int{}
	final := map[string][]string{}
	var used []string

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

		s = append(s, struct {
			name, lastName string
			score          float64
			departments    []string
		}{name, lastName, score, departments})
	}

	// Sort the students by score and first name
	sort.Slice(s, func(i, j int) bool {
		if s[i].score != s[j].score {
			return s[i].score > s[j].score
		}
		return s[i].name < s[j].name
	})

	for i := 0; i < 3; i++ {
		for _, entry := range s {
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

	// Print the output - Biotech first then Chem, Eng, Math, Physics
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
