package main

type params struct {
	duration      int
	intersections int
	streets       int
	routes        int
	score         int
}

type intersection struct {
	id         int
	inStreets  []*street
	outStreets []*street
}

type street struct {
	startIntersection *intersection
	endIntersection   *intersection
	name              string
	time              int
	semaphore		  *semaphore
}

type route struct {
	streets      []*street
	currentStreet int
	timeInStreet int
}

type semaphore struct {
	idStreet string
	street	 *street
	state    bool
	cars     []*route
	timePlan int
	timeOpen int
	historicOpen	[]int
}
