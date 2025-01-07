package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name           string
	Age            int
	Salary         float64
	EducationLevel string
}

type Statistics struct {
	AverageAge           float64
	YoungestPersons      []string
	OldestPersons        []string
	AverageSalary        float64
	HighestSalaryPersons []string
	LowestSalaryPersons  []string
	EducationLevelCounts map[string]int
}

func main() {
	inputFile := "people.json"
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	stats := calculateStats(people)

	outputFile := "statistics.json"

	writeToJson(stats, outputFile)
	fmt.Printf("Statistics written to %s\n", outputFile)

}

func calculateStats(people []Person) Statistics {
	var (
		totalAge               int
		totalSalary            float64
		youngestAge            = int(^uint(0) >> 1)
		oldestAge              = 0
		highestSalary          float64
		lowestSalary           = float64(^uint(0) >> 1)
		personByEducationLevel = make(map[string]int)
	)

	youngestPersons := []string{}
	oldestPersons := []string{}
	highestSalaryPersons := []string{}
	lowestSalaryPersons := []string{}

	for _, person := range people {

		totalAge += person.Age
		totalSalary += person.Salary

		if person.Age < youngestAge {
			youngestAge = person.Age
			youngestPersons = []string{person.Name}
		} else if person.Age == youngestAge {
			youngestPersons = append(youngestPersons, person.Name)
		}

		if person.Age > oldestAge {
			oldestAge = person.Age
			oldestPersons = []string{person.Name}
		} else if person.Age == oldestAge {
			oldestPersons = append(oldestPersons, person.Name)
		}

		if person.Salary > highestSalary {
			highestSalary = person.Salary
			highestSalaryPersons = []string{person.Name}
		} else if person.Salary == highestSalary {
			highestSalaryPersons = append(highestSalaryPersons, person.Name)
		}

		if person.Salary < lowestSalary {
			lowestSalary = person.Salary
			lowestSalaryPersons = []string{person.Name}
		} else if person.Salary == lowestSalary {
			lowestSalaryPersons = append(lowestSalaryPersons, person.Name)
		}

		personByEducationLevel[person.EducationLevel]++

	}

	averageAge := float64(totalAge) / float64(len(people))
	averageSalary := totalSalary / float64(len(people))

	return Statistics{
		AverageAge:           averageAge,
		YoungestPersons:      youngestPersons,
		OldestPersons:        oldestPersons,
		AverageSalary:        averageSalary,
		HighestSalaryPersons: highestSalaryPersons,
		LowestSalaryPersons:  lowestSalaryPersons,
		EducationLevelCounts: personByEducationLevel,
	}

}

func writeToJson(stats Statistics, outputFile string) {
	JsonData, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal statistics to JSON: %v", err)
	}

	err = os.WriteFile(outputFile, JsonData, 0644)
	if err != nil {
		log.Fatalf("Failed to write JSON to file: %v", err)
	}
}
