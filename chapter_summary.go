package main

import (
	"fmt"
	"unicode/utf8"
)

// iota-based grading scale
type Grade int

const (
	F Grade = iota // 0
	D              // 1
	C              // 2
	B              // 3
	A              // 4
)

// constants for eligibility check
const (
	MinAge = 18
	MaxAge = 65
)

func gradeName(g Grade) string {
	names := []string{"F", "D", "C", "B", "A"}
	return names[g]
}

func isEligible(age int) bool {
	return age >= MinAge && age <= MaxAge
}

func main() {
	// variable declarations - three styles
	var name string = "Kelly"
	var age = 21
	country := "Kenya"

	// zero values
	var score int
	var gpa float64
	var isGraduated bool
	fmt.Println("Zero values ->", score, gpa, isGraduated)

	// type conversion
	var totalUnits int = 24
	var maxUnits float64 = 30
	completion := float64(totalUnits) / maxUnits * 100
	fmt.Printf("Completion: %.1f%%\n", completion)

	// iota-based grade
	var currentGrade Grade = B
	fmt.Println("Current grade:", gradeName(currentGrade))

	// boolean logic with constants
	fmt.Println("Eligible for program:", isEligible(age))

	// byte vs rune vs len (handles Kelly's name if it had unicode, but here shown with country)
	sample := "Këlly"
	fmt.Println("len (bytes):", len(sample))
	fmt.Println("RuneCount (chars):", utf8.RuneCountInString(sample))

	// final formatted output
	fmt.Println("---- Student Profile ----")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Country:", country)
}