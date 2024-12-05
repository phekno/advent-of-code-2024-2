package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Print("error opening input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var line int
	numSafe := 0
	var intReports [][]int

	for scanner.Scan() {
		line++

		var report []int

		for _, strValue := range strings.Split(scanner.Text(), " ") {
			intValue, err := strconv.Atoi(strValue)
			if err != nil {
				log.Print("error converting to int")
			}

			report = append(report, intValue)
		}

		intReports = append(intReports, report)
	}

	for _, report := range intReports {
		if checkReportWithDeletion(report) {
			numSafe++
		} else {
			continue
		}
	}

	log.Printf("number of safe reports: %d\n", numSafe)
}

// /////////////////////////////////////////////////////////////////////////////
// Credit to https://www.bytesizego.com/blog/aoc-day2-golang because I could
// NOT get this working otherwise...shame on me.
// /////////////////////////////////////////////////////////////////////////////
func isSafe(report []int) bool {
	flagIncrease, flagDecrease := false, false

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if diff > 0 {
			flagIncrease = true
		} else if diff < 0 {
			flagDecrease = true
		} else {
			return false
		}

		if flagDecrease && flagIncrease {
			return false
		}

		if diff > 3 || diff < -3 {
			return false
		}
	}

	return true
}

func checkReportWithDeletion(report []int) bool {
	for i := 0; i < len(report); i++ {
		if isSafeWithDeletion(report, i) {
			return true
		}
	}

	return false
}

func isSafeWithDeletion(report []int, deleteIndex int) bool {
	copyReport := make([]int, len(report))
	copy(copyReport, report)

	if deleteIndex == len(copyReport)-1 {
		copyReport = copyReport[:deleteIndex]
	} else {
		copyReport = append(copyReport[:deleteIndex], copyReport[deleteIndex+1:]...)
	}

	return isSafe(copyReport)
}
