package main

import (
	"os"
	"strconv"
)

func getIntersectionNumberWithSemaphore(intersections *[]*intersection) int {
	i := 0
	for _, inter := range *intersections {
		if len(inter.inStreets) > 1 {
			i++
		}
	}
	return i
}

func exportResult(outputFile string, intersections *[]*intersection) {

    f, err := os.Create(outputFile)
    check(err)

    defer f.Close()

	// f.WriteString(strconv.Itoa(getIntersectionNumberWithSemaphore(intersections)) + "\n")
	f.WriteString(strconv.Itoa(len(*intersections)) + "\n")

	for i, inter := range *intersections {
		// if len(inter.inStreets) > 1 {	
			f.WriteString(strconv.Itoa(i) + "\n")
			f.WriteString(strconv.Itoa(len(inter.inStreets)) + "\n")
			for _,street := range inter.inStreets {
				// fmt.Println(street.semaphore)
				f.WriteString(street.name + " " + strconv.Itoa(getMinFromArray(&street.semaphore.historicOpen)) + "\n")
				// f.WriteString(street.name + " " + strconv.Itoa(getMaxFromArray(&street.semaphore.historicOpen)) + "\n")
				// f.WriteString(street.name + " " + strconv.Itoa(getAvgFromArray(&street.semaphore.historicOpen)) + "\n")
			}
		// }
	}
}

func getMinFromArray(array *[]int) int {
	if len(*array) == 0 {
		return 1
	}
	min := 500000000
	for _,i := range *array {
		if i < min {
			min = i
		}
	}
	return min
}

func getMaxFromArray(array *[]int) int {
	max := 1
	for _,i := range *array {
		if i > max {
			max = i
		}
	}
	return max
}

func getAvgFromArray(array *[]int) int {
	if len(*array) == 0 {
		return 1
	}
	sum := 0
	for _,i := range *array {
		sum += i
	}
	return sum / len(*array)
}