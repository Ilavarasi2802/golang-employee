package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Employee struct
type Employee struct {
	Id     int    `bson:"id"`
	Name   string `bson:"name"`
	Salary int    `bson:"salary"`
}

// Function that takes an array (fixed size) - Call by Value
func function1(arr [3]Employee) {
	fmt.Println("\nInside function1 array:")

	// Modify salaries (only inside function, won’t reflect in main)
	for i := range arr {
		arr[i].Salary += 1000
	}

	// Print inside function
	for _, emp := range arr {
		fmt.Printf("Id: %d, Name: %s, Salary: %d\n", emp.Id, emp.Name, emp.Salary)
	}
}

// Function that takes a slice (dynamic) - Call by Reference
func function2(slc []Employee) {
	fmt.Println("\nInside function2 slice:")

	// Modify salaries (will reflect in main)
	for i := range slc {
		slc[i].Salary += 2000
	}

	// Print inside function
	for _, emp := range slc {
		fmt.Printf("Id: %d, Name: %s, Salary: %d\n", emp.Id, emp.Name, emp.Salary)
	}
}

func main() {
	// ====================
	// 1. Create Array and Slice
	// ====================
	emp_array := [3]Employee{
		{Id: 101, Name: "Ilavarasi", Salary: 50000},
		{Id: 102, Name: "Sanjay Kumar", Salary: 60000},
		{Id: 103, Name: "Jaganath Mani", Salary: 70000},
	}

	emp_slice := []Employee{
		{Id: 201, Name: "Karthikeyan", Salary: 80000},
		{Id: 202, Name: "Srinithan", Salary: 90000},
		{Id: 203, Name: "Mahesh Babu", Salary: 100000},
	}

	// ====================
	// 2. Demonstrate Call by Value vs Call by Reference
	// ====================
	function1(emp_array)
	fmt.Println("\nBack in main after function1 array (No Change):")
	for _, emp := range emp_array {
		fmt.Printf("Id: %d, Name: %s, Salary: %d\n", emp.Id, emp.Name, emp.Salary)
	}

	function2(emp_slice)
	fmt.Println("\nBack in main after function2 slice (Changes Applied):")
	for _, emp := range emp_slice {
		fmt.Printf("Id: %d, Name: %s, Salary: %d\n", emp.Id, emp.Name, emp.Salary)
	}

	// ====================
	// 3. Connect to MongoDB
	// ====================
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := "mongodb+srv://ilavarasinataraj_db_user:RmkY5omt7xbLViA9@cluster0.d497idh.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("company").Collection("employees")

	// ====================
	// 4. Insert Array + Slice into MongoDB
	// ====================
	// Convert emp_array to slice first
	arrayDocs := make([]interface{}, len(emp_array))
	for i, emp := range emp_array {
		arrayDocs[i] = emp
	}

	// Convert emp_slice directly
	sliceDocs := make([]interface{}, len(emp_slice))
	for i, emp := range emp_slice {
		sliceDocs[i] = emp
	}

	// Insert both
	_, err = collection.InsertMany(ctx, arrayDocs)
	if err != nil {
		log.Fatal("Error inserting array employees:", err)
	}

	_, err = collection.InsertMany(ctx, sliceDocs)
	if err != nil {
		log.Fatal("Error inserting slice employees:", err)
	}

	fmt.Println("\nEmployees inserted into MongoDB successfully!")

	// ====================
	// 5. Retrieve Data from MongoDB
	// ====================
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal("Error retrieving employees:", err)
	}
	defer cursor.Close(ctx)

	var employees []Employee
	if err = cursor.All(ctx, &employees); err != nil {
		log.Fatal("Cursor error:", err)
	}

	fmt.Println("\nEmployees retrieved from MongoDB:")
	for _, emp := range employees {
		fmt.Printf("Id: %d, Name: %s, Salary: %d\n", emp.Id, emp.Name, emp.Salary)
	}
}
