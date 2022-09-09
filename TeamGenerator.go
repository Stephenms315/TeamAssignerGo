package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//----------Inputs---------
	names := [5]string{"a", "b", "d", "e", "f"}
	var grpSiz int = 2

	fmt.Println(names)
	fmt.Println(grpSiz)
	fmt.Println("-------")
	//----------Vars-----------
	
	//----------Spots----------
	//----------Assigner-------

}

func readMemberList() []string {
	//open student_names file
	file, err := os.Open("student_names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//Create Slice
	var students []string
	//Add student names from text file to slice
	for scanner.Scan() {
		students = append(students, scanner.Text())
	}

	return students
}

func randomizeAndAssign() {}

func printTeams() {}
