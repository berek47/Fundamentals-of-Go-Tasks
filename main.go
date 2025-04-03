package main

import (
	"fmt"
	"task1/grades"
)

func main() {
	var name string
	var numSubjects int

	// Get student name
	fmt.Print("Enter student name: ")
	fmt.Scanln(&name)

	// Get number of subjects
	fmt.Print("Enter the number of subjects: ")
	fmt.Scanln(&numSubjects)

	if numSubjects <= 0 {
		fmt.Println("Invalid number of subjects.")
		return
	}

	subjects := make(map[string]float64)
	var gradesList []float64

	for i := 0; i < numSubjects; i++ {
		var subject string
		var grade float64

		fmt.Printf("Enter subject %d name: ", i+1)
		fmt.Scanln(&subject)

		fmt.Printf("Enter grade for %s: ", subject)
		fmt.Scanln(&grade)

		if grade < 0 || grade > 100 {
			fmt.Println("Invalid grade. Please enter a grade between 0 and 100.")
			i--
			continue
		}

		subjects[subject] = grade
		gradesList = append(gradesList, grade)
	}

	// Calculate the average grade using the function from the grades package
	average := grades.CalculateAverage(gradesList)

	// Display results
	fmt.Printf("\nStudent Name: %s\n", name)
	fmt.Println("Subject Grades:")
	for subject, grade := range subjects {
		fmt.Printf("%s: %.2f\n", subject, grade)
	}
	fmt.Printf("Average Grade: %.2f\n", average)
}
