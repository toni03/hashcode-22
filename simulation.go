package main

func openSemaphoresInIntersections(intersections *[]*intersection) {
	for _, i := range *intersections {
		// if (len(i.inStreets) > 1) {
			var previousOpen *semaphore
			semaphoreToOpen := i.inStreets[0].semaphore
			startState := semaphoreToOpen.state
			for _, s := range i.inStreets {
				if len(semaphoreToOpen.cars) < len(s.semaphore.cars) {
					semaphoreToOpen = s.semaphore
					startState = semaphoreToOpen.state
				}
				if s.semaphore.state {
					previousOpen = s.semaphore
				}
				s.semaphore.state = false
			}
			semaphoreToOpen.state = true
			if startState == true {
				semaphoreToOpen.timeOpen++
			} else {
				semaphoreToOpen.timeOpen = 1
				if previousOpen != nil {
					previousOpen.historicOpen = append(previousOpen.historicOpen, previousOpen.timeOpen)
					previousOpen.timeOpen = 0
				}
			}
		// }
	}
}

func moveCarStreet(car *route) {
	car.currentStreet++
	if (len(car.streets) > car.currentStreet) {
		car.timeInStreet =  car.streets[car.currentStreet].time
	}
}

func updateSemaphoresStates(semaphores *[]*semaphore) {
	for _, s := range *semaphores {
		if s.state && len(s.cars) > 0 {
			moveCarStreet(s.cars[0])
			s.cars = s.cars[1:]
		}
	}
}

func updateCarsPosition(routes *[]*route) {
	for _,r := range *routes {
		if r.timeInStreet != 0 && r.currentStreet < len(r.streets) {
			r.timeInStreet--
			if r.timeInStreet == 0 {
				if r.streets[r.currentStreet].semaphore != nil {
					r.streets[r.currentStreet].semaphore.cars = append(r.streets[r.currentStreet].semaphore.cars, r)
				} else {
					moveCarStreet(r)
				}
			}
		}
	}
}

func simulation(simulationLength int, routes *[]*route, 
	intersections *[]*intersection, semaphores *[]*semaphore) {

	for i:=0; i < simulationLength; i++ {
		openSemaphoresInIntersections(intersections)
		// printSemaphores(semaphores)
		updateCarsPosition(routes)
		updateSemaphoresStates(semaphores)
	}

}