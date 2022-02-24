package main

import "fmt"

type person struct {
	name   string
	skills []skill
}

func (p *person) String() string {
	return fmt.Sprintf("Person{name: %s, skills: %v}", p.name, p.skills)
}

type project struct {
	name     string
	days     int
	score    int
	deadline int
	skills   []skill
}

func (p *project) String() string {
	return fmt.Sprintf("Project{name: %s, days: %d, score: %d, deadline: %d, skills: %v}", p.name, p.days, p.score, p.deadline, p.skills)
}

type skill struct {
	name  string
	level int
}

func (s *skill) String() string {
	return fmt.Sprintf("Skill{name: %s, level: %d}", s.name, s.level)
}
