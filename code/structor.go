package main

import "fmt"


type Student struct {
	ID int
	Grade int
	name string
	weight float64
}

func main(){


	var James Student = Student{1,10,"James", 68.9}
	var Allen Student = Student{2,50,"Allen",70.2}
	
	fmt.Println("######## Info of ", James.name, " ####################")
	fmt.Println("ID: ", James.ID)
	fmt.Println("Grade: ", James.Grade)
	fmt.Println("Name: ", James.name)
	fmt.Println("Weight: ", James.weight)
	fmt.Println("#######################################################\n")
	
	fmt.Println("######## Info of ", Allen.name, " ####################")
	fmt.Println("ID: ", Allen.ID)
	fmt.Println("Grade: ", Allen.Grade)
	fmt.Println("Name: ", Allen.name)
	fmt.Println("Weight: ", Allen.weight)
	fmt.Println("#######################################################")

}