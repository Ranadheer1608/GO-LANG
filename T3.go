package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Roll  int
	Marks map[string]float64
	Avg   float64
	Grade string
}

func (s *Student) Calculate() {
	var total float64
	for _, m := range s.Marks {
		total += m
	}
	s.Avg = total / float64(len(s.Marks))

	if s.Avg <= 50 {
		s.Grade = "Fail"
	} else if s.Avg > 50 && s.Avg < 70 {
		s.Grade = "B"
	} else {
		s.Grade = "A"
	}
}

func main() {
	students := []Student{
		{
			Name: "John",
			Roll: 101,
			Marks: map[string]float64{
				"Math":    85,
				"Science": 80,
				"English": 75,
			},
		},
		{
			Name: "Alex",
			Roll: 102,
			Marks: map[string]float64{
				"Math":    45,
				"Science": 50,
				"English": 40,
			},
		},
		{
			Name: "Bob",
			Roll: 103,
			Marks: map[string]float64{
				"Math":    60,
				"Science": 55,
				"English": 65,
			},
		},
	}

	results := make(map[int]Student)

	for _, s := range students {
		stu := s
		stu.Calculate()
		results[stu.Roll] = stu
	}

	fmt.Printf("%-5s %-10s %-10s %-6s\n", "Roll", "Name", "Average", "Grade")
	fmt.Println("------------------------------------------------")
	for _, s := range results {
		fmt.Printf("%-5d %-10s %-10.2f %-6s\n", s.Roll, s.Name, s.Avg, s.Grade)
	}
}
