package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func getParams() (string, string) {
	in := flag.String("in", "in/a.txt", "Input file")
	out := flag.String("out", "out/a.txt", "Output file")
	flag.Parse()
	return *in, *out
}

func main() {

	start := time.Now()

	// Parse file
	inputFile, outputFile := getParams()
	fmt.Printf("Input: %s Output: %s\n", inputFile, outputFile)

	f, err := os.Open(inputFile)
	check(err)
	reader := bufio.NewReader(f)
	data, _, err := reader.ReadLine()

	strs := strings.Split(string(data), " ")
	params := params{
		duration:      parseInt(strs[0]),
		intersections: parseInt(strs[1]),
		streets:       parseInt(strs[2]),
		routes:        parseInt(strs[3]),
		score:         parseInt(strs[4]),
	}
	fmt.Println("params")
	fmt.Println(params)

	intersections := make([]*intersection, params.intersections)
	for i := range intersections {
		intersections[i] = &intersection{
			id:         i,
			inStreets:  []*street{},
			outStreets: []*street{},
		}
	}

	streets := make(map[string]*street, params.streets)
	for i := 0; i < params.streets; i++ {
		data, _, err := reader.ReadLine()
		check(err)
		strs = strings.Split(string(data), " ")

		intersectionStart := intersections[parseInt(strs[0])]
		intersectionEnd := intersections[parseInt(strs[1])]

		st := street{
			startIntersection: intersectionStart,
			endIntersection:   intersectionEnd,
			name:              strs[2],
			time:              parseInt(strs[3]),
			semaphore:		   nil,
		}
		streets[strs[2]] = &st
		
		intersectionStart.outStreets = append(intersectionStart.outStreets, &st)
		intersectionEnd.inStreets = append(intersectionEnd.inStreets, &st)
	}

	semaphores := make([]*semaphore, 0)
	for _, v := range intersections {
		// if len(v.inStreets) > 1 {
		for _, v := range v.inStreets {
			sem := semaphore{
				idStreet: v.name,
				street:   v,
				state:    false,
				cars:     []*route{},
			}
			semaphores = append(semaphores, &sem)
			v.semaphore = &sem
		}
		// }
	}

	routes := make([]*route, params.routes)
	for i := 0; i < params.routes; i++ {
		data, _, err := reader.ReadLine()
		check(err)
		strs = strings.Split(string(data), " ")[1:]
		routeStreets := make([]*street, len(strs))
		for j, s := range strs {
			st := streets[s]
			routeStreets[j] = st
		}

		r := route{
			streets:      routeStreets,
			currentStreet: 0,
			timeInStreet: 1,
		}
		routes[i] = &r

		if routeStreets[0].semaphore != nil {
			routeStreets[0].semaphore.cars = append(routeStreets[0].semaphore.cars, &r)
		}
	}

	timeLoad := time.Since(start)
	fmt.Printf("Time loading: %v, Streets: %v, Routes: %v, Intersections: %v, Semaphores: %v.", 
		timeLoad, len(streets), len(routes), len(intersections), len(semaphores))
	fmt.Println();
	
	startSimulation := time.Now()
	simulation(params.duration, &routes, &intersections, &semaphores)
	// printSemaphores(&semaphores)
	timeSimulation := time.Since(startSimulation)
	fmt.Printf("Time simulation: %v", timeSimulation)
	fmt.Println();
	
	startExport := time.Now()
	exportResult(outputFile, &intersections)
	timeExport := time.Since(startExport)
	fmt.Printf("Time export: %v", timeExport)
	fmt.Println();
}

func printSemaphores(semaphores *[]*semaphore) {
	for _, s := range *semaphores {
		fmt.Printf("Semaphore street: %v, state: %v, heapCars: %v, historicChanges: %v, timeOpen: %v",
			s.idStreet, s.state, len(s.cars), s.historicOpen, s.timeOpen)
		fmt.Println()
	}
}
