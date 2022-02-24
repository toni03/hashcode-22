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

var personsBySkil map[string][]person
var projectByPeople map[string]string

func main() {
	projectByPeople = map[string]string{}
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

	personsBySkill := indexPersons(persons)
	fmt.Println("indexedPersons", personsBySkill)

	projectsDone := 0
	for {
		fmt.Println("projects", projects)
		if projectsDone == len(projects) {
			return
		}

		for _, project := range projects {
			if project.isDone {
				continue
			}
			if project.onGoing() {
				project.remainingDays -= 1
				if project.remainingDays == 0 {
					unAssignPeopleToProject(project)
					project.isDone = true
					projectsDone += 1
				}
				continue
			}

			personCandidates := getCandidatesForProject(project)
			fmt.Println(personCandidates)
			if len(personCandidates) == len(project.skills) {
				assignPeopleToProject(project, personCandidates)
			}
		}
	}
}

func getCandidatesForProject(project project) []person {
	people := []person{}
	for _, skill := range project.skills {
		for _, person := range personsBySkil[skill.name] {
			if person.indexedSkills[skill.name] >= skill.level {
				// TODO: corner cases
				if isPersonFree(person) {
					people = append(people, person)
				}
			}
		}
	}
	return people
}

func assignPeopleToProject(project project, people []person) {
	for _, p := range people {
		projectByPeople[p.name] = project.name
	}
	project.remainingDays = project.days
	project.people = people
}

func unAssignPeopleToProject(project project) {
	for _, p := range project.people {
		delete(projectByPeople, p.name)
	}
}

func isPersonFree(p person) bool {
	_, ok := projectByPeople[p.name]
	return !ok
}
