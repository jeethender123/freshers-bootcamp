//package exercise3

package main

import (
	"fmt"
)

type Salary interface {
	compute_salary() int
}

type fullTime struct {
	duration int
}

type contractor struct {
	duration int
}

type freelancer struct {
	duration int
}


func (employee fullTime) compute_salary() int {
	return employee.duration*500
}

func (employee contractor) compute_salary() int {
	return employee.duration*100
}

func (employee freelancer) compute_salary() int {
	return employee.duration*10
}


func main() {
	fte:=fullTime{30}
	cont:=contractor{20}
	fl:=freelancer{250}
	fmt.Println(fte.compute_salary())
	fmt.Println(cont.compute_salary())
	fmt.Println(fl.compute_salary())
}

