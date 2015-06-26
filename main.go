package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"strings"
)

func main() {
	excelFileName := os.Args[1]
	xlFile, _ := xlsx.OpenFile(excelFileName)

	students := []*Student{}
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

	// matching
	matches := []*Match{}
	matchByStudent := map[*Student]map[int]*Match{}
	for n, s1 := range students {
		for _, s2 := range students[n+1:] {
			match := &Match{match(s1, s2), s1, s2}
			matches = append(matches, match)
			for _, student := range []*Student{s1, s2} {
				_, found := matchByStudent[student]
				if !found {
					matchByStudent[student] = map[int]*Match{}
				}
				matchByStudent[student][match.Score] = match
			}
		}
	}
	fmt.Println(matches)
	fmt.Println(matchByStudent)
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

type Match struct {
	Score              int
	Student1, Student2 *Student
}

func (m Match) String() string {
	return fmt.Sprintf("match between %s and %s with a score of %d\n", m.Student1, m.Student2, m.Score)
}

func match(s1, s2 *Student) int {
	score := 0

	// Check if there is a common language.
	for language := range s1.SpokenLanguages {
		_, contains := s2.SpokenLanguages[language]
		if contains {
			score += 1
		}
	}

	// Check if they share the same study.
	if s1.Study == s2.Study {
		score += 1
	}

	return score
}

type Gender string

const (
	M Gender = "M"
	F        = "F"
)
