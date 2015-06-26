package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"strings"
)

func main() {
	excelFileName := os.Args[1]
	xlFile, err := xlsx.OpenFile(excelFileName)

	students := []*Student{}

	if err != nil {
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			name := row.Cells[0].Value
			sex := Gender(row.Cells[1].Value)
			languages := map[string]struct{}{}
			for _, language := range strings.Split(row.Cells[2].Value, ",") {
				languages[language] = struct{}{}
			}
			study := row.Cells[3].Value
			student := &Student{name, sex, languages, study}
			students = append(students, student)
		}
	}

	fmt.Println(students)

	// matching
	for _, s1 := range students {
		for _, s2 := range students {
			if s1 == s2 {
				continue
			}
			fmt.Printf("%s and %s have a match of %.1f\n", s1.Name, s2.Name, match(s1, s2))
		}
	}

}

type Student struct {
	Name            string
	Sex             Gender
	SpokenLanguages map[string]struct{}
	Study           string
}

func (s Student) String() string {
	return s.Name
}

func match(s1, s2 *Student) float64 {
	score := 0.0

	// Check if there is a common language.
	for language := range s1.SpokenLanguages {
		_, contains := s2.SpokenLanguages[language]
		if contains {
			score += 0.5
		}
	}

	// Check if they share the same study
	if s1.Study == s2.Study {
		score += 0.5
	}

	return score
}

type Gender string

const (
	M Gender = "M"
	F        = "F"
)
