package main


import "fmt"

type getSalary interface{
	FullTimeEmp() int
	Contractor() int
	Freelancer() int
}
type Employees struct{
	FullTime_days, Contractor_days, Freelancer_hours int
}
func (e Employees) FullTimeEmp() int {
	return e.FullTime_days*500
}
func (e Employees) Contractor() int {
	return e.Contractor_days*100
}
func (e Employees) Freelancer() int {
	return e.Freelancer_hours*10
}

func main(){
	var employee getSalary
	var workingdays_Ft,workingdays_cont,workinghours_free int
	fmt.Println("enter values:\n")
	fmt.Scanf("%d",&workingdays_Ft)
	fmt.Scanf("%d",&workingdays_cont)
	fmt.Scanf("%d",&workinghours_free)
	employee = Employees{workingdays_Ft,workingdays_cont,workinghours_free}
	fmt.Println("FT employee salary: ", employee.FullTimeEmp())
	fmt.Println("Contractor salary: ", employee.Contractor())
	fmt.Println("Freelancer Salary: ",employee.Freelancer())
}
