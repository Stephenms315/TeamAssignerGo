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
	var numStu int = len(names)
	var numGrp int = (numStu / grpSiz)
	var grpRm int = numStu - (grpSiz * numGrp)

	fmt.Println(numStu)
	fmt.Println(numGrp)
	fmt.Println(grpSiz)
	fmt.Println(grpRm)
	fmt.Println("-------")
	//----------Spots----------
	grpRang := makeRange(1, numGrp)
	grpsL := []int{}

	for i := 1; i < numGrp+1; i++ {
		grpsL = append(grpsL, grpRang...)
	}
	fmt.Println(grpsL)
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

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func randomizeAndAssign() {}

func printTeams() {}
