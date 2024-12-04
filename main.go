package main

import (
	"bufio"
	"log"
	"math"
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
	var totalReport []bool

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
		reportIsSafe, err := isSafe(report)
		if err != nil {
			log.Print("error checking report for safety")
		}

		totalReport = append(totalReport, reportIsSafe)
	}

	for _, report := range totalReport {
		if report {
			numSafe++
		}
	}

	log.Printf("number of safe reports: %d\n", numSafe)
}

func isSafe(report []int) (bool, error) {
	var increasing []bool
	var distGood []bool
	allIncreasing := 0
	allDistGood := 0

	for i, value := range report {
		if i >= len(report)-1 {
			break
		}
		currValue := value
		nextValue := report[i+1]

		difference := int(math.Abs(float64(nextValue - currValue)))

		if difference >= 1 && difference <= 3 {
			distGood = append(distGood, true)
		} else {
			distGood = append(distGood, false)
		}

		if nextValue > currValue {
			increasing = append(increasing, true)
		} else {
			increasing = append(increasing, false)
		}
	}

	for i, inc := range increasing {
		if i >= len(increasing)-1 {
			break
		}
		if inc == increasing[i+1] {
			continue
		} else {
			allIncreasing++
		}
	}

	for i, dist := range distGood {
		if i >= len(distGood)-1 {
			break
		}
		if dist == distGood[i+1] {
			continue
		} else {
			allDistGood++
		}
	}

	if allDistGood == 0 && allIncreasing == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
