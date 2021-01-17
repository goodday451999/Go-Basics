package main

import (
	"fmt"
)

type student struct {
	name string
	grades []int
	age int
} 

func (s student) getAge() {
	fmt.Println(s.age)
} 

func (s *student) setAge(a int) {
	s.age = a
}

func (s *student) updateGrades() {
	s.grades = append(s.grades, 46)
}

func (s *student) getAvgGrades() float64 {
	sum := 0
	for _, num := range s.grades {
		sum += num
 	}
	return float64(sum) / float64(len(s.grades))
}
	
 
func main(){
	s := student{"abc xyz", []int{10, 20}, 21}
	fmt.Println(s) //{abc xyz [10 20] 21}
	s.getAge() // 21	
	
	s.setAge(22)	
	 
	fmt.Println(s) // {abc xyz [10 20] 21} for pass by val
		       // {abc xyz [10 20] 22} for pass by ref
		
	s1 := s	 
	fmt.Println(s1)  // {abc xyz [10 20] 22} same as above
	s1.updateGrades()
	fmt.Println(s1)	// {abc xyz [10 20 40] 22}
		
	avg := s.getAvgGrades()
	avg1 := s1.getAvgGrades()
	fmt.Println(avg, avg1) // 15 25.333333333333332
}

