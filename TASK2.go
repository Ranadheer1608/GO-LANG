package main

import (
	"fmt"
)

func getGrade(mark int) string {
	switch {
	case mark >= 90:
		return "Grade A"
	case mark >= 70:
		return "Grade B"
	case mark >= 50:
		return "Grade C"
	default:
		return "Fail"
	}
}

func main() {
	originalMarks := []int{95, 82, 47, 68, 73, 38, 90, 59, 49, 41}

	reExamMarks := []int{0, 0, 55, 0, 0, 52, 0, 0, 58, 50}

	passCount := 0
	failCount := 0

	fmt.Printf("%-10s %-10s %-20s %-10s\n", "Marks", "Grade", "Pass Code", "Final Marks")

	for i := 0; i < len(originalMarks); i++ {
		origMark := originalMarks[i]
		reExamMark := reExamMarks[i]
		finalMark := origMark
		passCode := ""

		if origMark >= 50 {
			passCode = "pass"
		} else {
			if reExamMark >= 50 {
				finalMark = reExamMark
				passCode = "passed"
			} else {
				finalMark = origMark
				passCode = "fail"
			}
		}

		grade := getGrade(finalMark)

		if finalMark >= 50 {
			passCount++
		} else {
			failCount++
		}

		fmt.Printf("%-10d %-10s %-20s %-10d\n", origMark, grade, passCode, finalMark)
	}

	fmt.Printf("\nTotal Passed: %d\n", passCount)
	fmt.Printf("Total Failed: %d\n", failCount)
}
