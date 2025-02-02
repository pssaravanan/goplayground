package gotutorial

import "fmt"

type Person interface {
	printName()
}

type Employee struct {
	name string
	dept string
}

func (e Employee) printName(){
	fmt.Println("\t Employee Name: ", e.name)
}

type Student struct{
	name string
	m1 int
	m2 int
	m3 int
}

func (s Student) printName() {
	fmt.Println("\t Student Name:", s.name)
}

func (s Student) total() int{
	return s.m1 + s.m2 + s.m3
}

func (s *Student) addGraceMark(){
	s.m1 = s.m1 + 10
	s.m2 = s.m2 + 5
	s.m3 = s.m3 + 7
}

func MethodMain(){
	var s = Student{"Saravanan", 98, 95, 89}
	fmt.Println("Methods:")
	s.printName()
	fmt.Println("\t Total:", s.total())
	s.addGraceMark()
	fmt.Println("\t Total after Grace mark:", s.total())

	var p Person
	p = Employee{name: "Saravanan", dept: "Engineering"}
	p.printName()
	p = s
	p.printName()

	var nilStudent Student
	p = nilStudent
	nilStudent.printName()
}