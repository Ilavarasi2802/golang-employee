package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary int
}

// Function that takes an array (fixed size)
func function1(arr [3]Employee) {
	fmt.Println("\nInside function1 array:")

	// Modify salaries
	for i := range arr {
		arr[i].Salary += 1000
	}

	// Print inside function
	for _, emp := range arr {
		fmt.Printf("Id: %d, Name: %s, Salary: %d\n", emp.Id, emp.Name, emp.Salary)
	}
}

// Function that takes a slice (dynamic)
func function2(slc []Employee) {
	fmt.Println("\nInside function2 slice:")

	// Modify salaries
	for i := range slc {
		slc[i].Salary += 2000
	}

	// Print inside function
	for _, emp := range slc {
		fmt.Printf("Id: %d, Name: %s, Salary: %d\n", emp.Id, emp.Name, emp.Salary)
	}
}

func main() {
	// Array (fixed length = 3)
	emp_array := [3]Employee{
		{Id: 101, Name: "Ilavarasi", Salary: 50000},
		{Id: 102, Name: "Sanjay Kumar", Salary: 60000},
		{Id: 103, Name: "Jaganath Mani", Salary: 70000},
	}

	// Slice (dynamic length)
	emp_slice := []Employee{
		{Id: 201, Name: "Karthikeyan", Salary: 80000},
		{Id: 202, Name: "Srinithan", Salary: 90000},
		{Id: 203, Name: "Mahesh Babu", Salary: 100000},
	}

	// Call function1 (array - (Call by Value))
	function1(emp_array)
	fmt.Println("\nBack in main after function1 array:")
	for _, emp := range emp_array {
		fmt.Printf("Id: %d, Name: %s, Salary: %d\n", emp.Id, emp.Name, emp.Salary)
	}

	// Call function2 (slice - (Call by Reference))
	function2(emp_slice)
	fmt.Println("\nBack in main after function2 slice:")
	for _, emp := range emp_slice {
		fmt.Printf("Id: %d, Name: %s, Salary: %d\n", emp.Id, emp.Name, emp.Salary)
	}
}
