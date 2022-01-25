package main

import (
	"fmt"
	"github.com/bluebook/golang/module/magazine"
)

func main() {
	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	fmt.println(employee.Name)
	fmt.println(employee.Salary)
}