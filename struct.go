package main

import "fmt"

type Employee struct {
	Emp_Id     int
	Emp_Name   string
	Emp_Salary int
}

func main() {
	var count int

	// Ask how many employees to add
	fmt.Print("Enter number of employees: ")
	fmt.Scanln(&count)

	// Slice to store multiple employees
	var employees []Employee

	// Let's add employees using a loop
	for i := 1; i <= count; i++ {
		var id int
		var name string
		var salary int

		fmt.Printf("Enter details for Employee %d\n", i)
		fmt.Print("Enter Employee Id: ")
		fmt.Scan(&id)
		fmt.Print("Enter Employee Name: ")
		fmt.Scan(&name)
		fmt.Print("Enter Employee Salary: ")
		fmt.Scan(&salary)

		// Create employee and add to slice
		emp := Employee{Emp_Id: id, Emp_Name: name, Emp_Salary: salary}
		employees = append(employees, emp)
	}

	// Display all employees
	fmt.Println("\nEmployee Info")
	for _, emp := range employees {
		fmt.Printf("Id: %d, Name: %s, Salary: %.d\n", emp.Emp_Id, emp.Emp_Name, emp.Emp_Salary)
	}
}
