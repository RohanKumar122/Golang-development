package main

import "fmt"

type Person struct {
	firstNname string
	lastName   string
	age        int
}

type Contact struct {
	email string
	phone string
}

type Address struct {
	house string
}

type Employee struct {
	person_details Person
	person_contact Contact
	person_address Address
}

func main() {
	var rohan Person
	fmt.Println(rohan)
	rohan.firstNname = "Rohan"
	rohan.lastName = "Kumar"
	rohan.age = 24
	fmt.Println(rohan)

	// Method 2: Using new keyword
	var Person2 = new(Person)
	Person2.firstNname = "Akash"
	Person2.lastName = "Kumar"
	Person2.age = 24
	println(Person2)
	fmt.Println(Person2)

	// Employee
	employee1 := new(Employee)
	employee1.person_details = Person{
		firstNname: "Rohan",
		lastName:   "Kumar",
		age:        24,
	}
	employee1.person_contact = Contact{
        email: "rohan@emse.uk",
		phone:"92334352522",
	}
	employee1.person_address.house="3rd floor house"

	fmt.Println("The employee details: ",employee1.person_details)
	fmt.Println("Complete employee details: ",employee1)

}
