package main

import "fmt"

type person struct {
	name          string
	skills        []skill
	indexedSkills map[string]int
}

func (p *person) indexSkills() {
	index := map[string]int{}
	for _, s := range p.skills {
		index[s.name] = s.level
	}
	p.indexedSkills = index
}

func (p person) String() string {
	return fmt.Sprintf("Person{name: %s, skills: %v}", p.name, p.skills)
}

type project struct {
	name     string
	days     int
	score    int
	deadline int

	skills        []skill
	remainingDays int
	isDone        bool
	people        []person
	indexedSkills map[string]int
}

func (p project) onGoing() bool {
	return p.remainingDays > 0
}

func (p project) indexSkills() {
	index := map[string]int{}
	for _, s := range p.skills {
		index[s.name] = s.level
	}
	p.indexedSkills = index
}

func (p project) String() string {
	return fmt.Sprintf("Project{name: %s, days: %d, score: %d, deadline: %d, skills: %v, people: %v}", p.name, p.days, p.score, p.deadline, p.skills, p.people)
}

type skill struct {
	name  string
	level int
}

func (s skill) String() string {
	return fmt.Sprintf("Skill{name: %s, level: %d}", s.name, s.level)
}
