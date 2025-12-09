package main

import "fmt"

type student struct {
	Name   string
	MatNr  int
	Grades []int
	Active bool
}

func NewStudent(name string, matrNr int) student {
	return student{
		Name:   name,
		MatNr:  matrNr,
		Grades: []int{},
		Active: true,
	}
}

func (s *student) AddGrade(grade int) {
	s.Grades = append(s.Grades, grade)

}

func (s student) GetAverage() float64 { //gleich rekursiv probieren
	Summe := 0
	if len(s.Grades) == 0 {
		return 0.0
	}

	for _, el := range s.Grades {
		Summe += el
	}

	return float64(Summe) / float64(len(s.Grades))

}

func (s *student) Exmatrikulieren() {
	s.Active = false
}

func ExampleNewStudent() {

	// 1. Wir rufen die FUNKTION auf (kein Receiver!)
	// Wir werfen Daten (Name, MatNr) oben rein...
	neuerStudent := NewStudent("Erika Musterfrau", 99999)

	// ... und es kommt ein fertiges Struct unten raus.

	// 2. Wir prüfen, ob alles stimmt:
	fmt.Println(neuerStudent.Name)
	fmt.Println(neuerStudent.MatNr)
	fmt.Println(neuerStudent.Active)
	fmt.Println(neuerStudent.Grades)

	// Das hier prüft Go automatisch:
	// Output:
	// Erika Musterfrau
	// 99999
	// true
	// []
}
