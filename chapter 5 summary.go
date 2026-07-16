package main

import "fmt"

// classify takes a numeric score and returns a letter grade.
// This uses a "switch with no condition" - a clean alternative
// to writing a long if/else-if chain.
func classify(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func main() {
	// A slice (ordered list) of exam scores to analyze.
	scores := []int{95, 42, 78, 88, -5, 67, 105, 55, 90, 30}

	total := 0       // running sum of valid scores
	countValid := 0  // how many scores were actually valid
	countInvalid := 0 // how many scores were skipped

	fmt.Println("---- Processing Scores ----")

	// range lets us loop over a slice, giving us both the index
	// and the value at each position. We use _ to discard the index
	// here since we don't need it.
	for _, score := range scores {

		// --- Validation using if/else ---
		// A valid exam score should be between 0 and 100.
		// Anything outside that range is bad data (e.g. -5 or 105),
		// so we skip it using 'continue', which jumps straight to
		// the next loop iteration without running the rest of the code below.
		if score < 0 || score > 100 {
			fmt.Println("Skipping invalid score:", score)
			countInvalid++
			continue
		}

		// If we reach this point, the score is valid.
		countValid++
		total += score

		// --- Classify the valid score using switch ---
		grade := classify(score)
		fmt.Printf("Score: %d -> Grade: %s\n", score, grade)

		// --- Early exit example using break ---
		// If we hit a perfect run of 3 valid A's in a row conceptually,
		// we could stop - here we demonstrate break by stopping
		// early if we encounter a 100 (a perfect score).
		if score == 100 {
			fmt.Println("Perfect score hit! Stopping analysis early.")
			break
		}
	}

	fmt.Println("\n---- Summary Report ----")
	fmt.Println("Valid scores processed:", countValid)
	fmt.Println("Invalid scores skipped:", countInvalid)

	// --- Classic three-part for loop ---
	// Just to demonstrate the "traditional" for loop style,
	// we print a countdown before showing the final average.
	fmt.Println("\nFinalizing report:")
	for i := 3; i > 0; i-- {
		fmt.Println(i, "...")
	}

	// --- Average calculation with type conversion ---
	// total and countValid are both int, but we want a decimal
	// average, so we convert to float64 before dividing.
	if countValid > 0 {
		average := float64(total) / float64(countValid)
		fmt.Printf("Average score: %.2f\n", average)

		// --- Nested if/else for final remark ---
		if average >= 75 {
			fmt.Println("Overall performance: Good")
		} else if average >= 50 {
			fmt.Println("Overall performance: Average")
		} else {
			fmt.Println("Overall performance: Needs improvement")
		}
	} else {
		// This handles the edge case where every score was invalid
		fmt.Println("No valid scores to calculate an average.")
	}
}