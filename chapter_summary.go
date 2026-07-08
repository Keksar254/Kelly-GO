package main

import (
	"fmt"
	"unicode/utf8"
)


type Grade int

const (
	F Grade = iota 
	D              
	C              
	B              
	A              
)


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
	
	var name string = "Kelly"
	var age = 21
	country := "Kenya"

	
	var score int
	var gpa float64
	var isGraduated bool
	fmt.Println("Zero values ->", score, gpa, isGraduated)

	
	var totalUnits int = 24
	var maxUnits float64 = 30
	completion := float64(totalUnits) / maxUnits * 100
	fmt.Printf("Completion: %.1f%%\n", completion)


	var currentGrade Grade = B
	fmt.Println("Current grade:", gradeName(currentGrade))

	
	fmt.Println("Eligible for program:", isEligible(age))


	sample := "Këlly"
	fmt.Println("len (bytes):", len(sample))
	fmt.Println("RuneCount (chars):", utf8.RuneCountInString(sample))

	
	fmt.Println("---- Student Profile ----")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Country:", country)
}