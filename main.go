package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func getParams() (string, string) {
	in := flag.String("in", "in/a_an_example.in.txt", "Input file")
	out := flag.String("out", "out/a.txt", "Output file")
	flag.Parse()
	return *in, *out
}

func main() {
	// Parse file
	inputFile, outputFile := getParams()
	fmt.Printf("Input: %s Output: %s\n", inputFile, outputFile)

	f, err := os.Open(inputFile)
	check(err)
	reader := bufio.NewReader(f)
	data, _, err := reader.ReadLine()
	check(err)

	strs := strings.Split(string(data), " ")
	numPersons := parseInt(strs[0])
	numProjects := parseInt(strs[1])

	persons := make([]person, numPersons)

	for i := 0; i < numPersons; i++ {
		data, _, err := reader.ReadLine()
		check(err)
		strs = strings.Split(string(data), " ")

		persons[i].name = strs[0]
		numSkills := parseInt(strs[1])
		persons[i].skills = make([]skill, numSkills)
		for j := 0; j < numSkills; j++ {
			data, _, err := reader.ReadLine()
			check(err)
			strs = strings.Split(string(data), " ")

			persons[i].skills[j].name = strs[0]
			persons[i].skills[j].level = parseInt(strs[1])
		}
	}

	projects := make([]project, numProjects)
	for i := 0; i < numProjects; i++ {
		data, _, err := reader.ReadLine()
		check(err)
		strs = strings.Split(string(data), " ")

		projects[i].name = strs[0]
		projects[i].days = parseInt(strs[1])
		projects[i].score = parseInt(strs[2])
		projects[i].deadline = parseInt(strs[3])

		numSkills := parseInt(strs[4])
		projects[i].skills = make([]skill, numSkills)

		for j := 0; j < numSkills; j++ {
			data, _, err := reader.ReadLine()
			check(err)
			strs = strings.Split(string(data), " ")

			projects[i].skills[j].name = strs[0]
			projects[i].skills[j].level = parseInt(strs[1])
		}
	}

	fmt.Println("persons", persons)
	fmt.Println("projects", projects)
}
