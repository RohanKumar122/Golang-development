package main

import "fmt"

func main() {
	studentsGrades := make(map[string]int)

	studentsGrades["Rohan"] = 80
	studentsGrades["Akash"] = 70
	studentsGrades["Zebra"] = 60
	studentsGrades["Rio"] = 90
	studentsGrades["Aman"] = 44

	fmt.Println("Grade of Rohan is ", studentsGrades["Rohan"])
	fmt.Println("Grade of Bob is ", studentsGrades["Bob"]) //if not exist it returns 0
	delete(studentsGrades,"Akash")

	// checking if key exists
	grades, exists := studentsGrades["dravid"]
	fmt.Println("Dradesof Dravid is: ", grades)
	fmt.Println("Dravid exists: ", exists)

	// checking if key exists
	Grades,Exists := studentsGrades["Rohan"]
	fmt.Println("Grades  of Rohan is: ",Grades)
	fmt.Println("Rohan Exist : ",Exists)

	for index,value :=range studentsGrades{
		fmt.Printf("---------key is %s and value is %d\n", index,value)
	}

}
